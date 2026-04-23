const CACHE_NAME = "doprava-v1";
const ASSETS = [
	"/",
	"/public/manifest.webmanifest",
	"/public/icon/android-icon-192x192.png",
	"/public/icon/android-icon-512x512.png"
];

self.addEventListener("install", event => {
	event.waitUntil(
		caches.open(CACHE_NAME).then(cache => cache.addAll(ASSETS))
	);
});

self.addEventListener("activate", event => {
	event.waitUntil(
		caches.keys().then(keys =>
			Promise.all(
				keys
					.filter(key => key !== CACHE_NAME)
					.map(key => caches.delete(key))
			)
		)
	);
});

self.addEventListener("fetch", event => {
	if (event.request.method !== "GET") return;

	event.respondWith(
		caches.match(event.request).then(cached => {
			return cached || fetch(event.request);
		})
	);
});