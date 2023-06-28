(async () => {
    if ('serviceWorker' in navigator) {
        try {
            console.log(`开始注册 ServiceWorker...`);
            await navigator.serviceWorker.register(
                'worker.js' + Date.now(),
                {
                    scope: './',
                }
            );
            console.log(`ServiceWorker 注册成功`);
        } catch (error) {
            console.error(`ServiceWorker 注册失败：${error}`);
        }
    }
})()