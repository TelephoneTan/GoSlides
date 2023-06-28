self.addEventListener('install', () => {
    console.log(`ServiceWorker 安装中...`);
    self.skipWaiting();
});

function play(urls, i) {
    if (i === urls.length) {
        i = 0
    }
    clients.matchAll({type: "window"})
        .then((clientList) => {
            for (const client of clientList) {
                client.navigate(urls[i])
            }
            setTimeout(() => {
                play(urls, ++i)
            }, 3500);
        })

}

self.addEventListener('activate', (event) => {
    console.log(`ServiceWorker 激活中...`);
    play([/*INSERT URLS*/], 0)
    event.waitUntil(new Promise(resolve =>
        resolve((async () => {
            await clients.claim()
        })())
    ).then(() => console.log(`ServiceWorker 已激活`)))
});