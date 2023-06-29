self.addEventListener('install', () => {
    console.log(`ServiceWorker 安装中...`);
    self.skipWaiting();
});

const urls = [/*INSERT URLS*/]

let i = 0

async function play() {
    if (i === urls.length) {
        i = 0
    }
    const cls = await clients.matchAll({type: "window"})
    for (const client of cls) {
        await client.navigate(urls[i])
    }
    i++
}

const taskNext = "next"

self.addEventListener("periodicsync", (event) => {
    if (event.tag === taskNext) {
        event.waitUntil(play());
    }
});

const period = 3500

self.addEventListener('activate', (event) => {
    console.log(`ServiceWorker 激活中...`);
    event.waitUntil(new Promise(resolve =>
        resolve((async () => {
            await clients.claim()
            try {
                await registration.periodicSync.register(taskNext, {
                    minInterval: period,
                });
            } catch (error) {
                console.error(`定时任务创建失败：${error}`);
                setInterval(play, period)
                await play()
            }
        })())
    ).finally(() => console.log(`ServiceWorker 已激活`)))
});