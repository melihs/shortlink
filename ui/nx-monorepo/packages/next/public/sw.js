if(!self.define){let e,s={};const a=(a,n)=>(a=new URL(a+".js",n).href,s[a]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=a,e.onload=s,document.head.appendChild(e)}else e=a,importScripts(a),s()})).then((()=>{let e=s[a];if(!e)throw new Error(`Module ${a} didn’t register its module`);return e})));self.define=(n,c)=>{const t=e||("document"in self?document.currentScript.src:"")||location.href;if(s[t])return;let i={};const r=e=>a(e,t),d={module:{uri:t},exports:i,require:r};s[t]=Promise.all(n.map((e=>d[e]||r(e)))).then((e=>(c(...e),i)))}}define(["./workbox-1051b61c"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/next/_next/static/chunks/242-cf6360725c9a47c3.js",revision:"cf6360725c9a47c3"},{url:"/next/_next/static/chunks/242-cf6360725c9a47c3.js.map",revision:"d157db0d7bc5b216bdc3719b20660a03"},{url:"/next/_next/static/chunks/252-87a8446e526ff70c.js",revision:"87a8446e526ff70c"},{url:"/next/_next/static/chunks/252-87a8446e526ff70c.js.map",revision:"630c334974b57e542bfc4ddbce344e1b"},{url:"/next/_next/static/chunks/2edb282b-99fa8a592c3e6564.js",revision:"99fa8a592c3e6564"},{url:"/next/_next/static/chunks/2edb282b-99fa8a592c3e6564.js.map",revision:"e8bfcc3c2d9e8318e67cca20b7cf968e"},{url:"/next/_next/static/chunks/330-316387a756051cb7.js",revision:"316387a756051cb7"},{url:"/next/_next/static/chunks/330-316387a756051cb7.js.map",revision:"ef553ab00dcbc2288127829b20f6bd5f"},{url:"/next/_next/static/chunks/432-17c17ba52b93b09e.js",revision:"17c17ba52b93b09e"},{url:"/next/_next/static/chunks/432-17c17ba52b93b09e.js.map",revision:"8b63eec7d4777e2724630af9d87d5461"},{url:"/next/_next/static/chunks/579-351c62aa2ca91c36.js",revision:"351c62aa2ca91c36"},{url:"/next/_next/static/chunks/579-351c62aa2ca91c36.js.map",revision:"55532229f29b656786ba1d50c42ca736"},{url:"/next/_next/static/chunks/685-4e7c6149d2c36f0b.js",revision:"4e7c6149d2c36f0b"},{url:"/next/_next/static/chunks/685-4e7c6149d2c36f0b.js.map",revision:"da619416c6f336a1b09c4aabb674ebd6"},{url:"/next/_next/static/chunks/690-809f0a4e065b9648.js",revision:"809f0a4e065b9648"},{url:"/next/_next/static/chunks/690-809f0a4e065b9648.js.map",revision:"73d7df0dc1d6ab08e3d2ec482716992d"},{url:"/next/_next/static/chunks/706-26878f53b8f0222e.js",revision:"26878f53b8f0222e"},{url:"/next/_next/static/chunks/706-26878f53b8f0222e.js.map",revision:"c621a572e77412d01681eac12b94bb4f"},{url:"/next/_next/static/chunks/75-3c6a0e70c72c3b85.js",revision:"3c6a0e70c72c3b85"},{url:"/next/_next/static/chunks/75-3c6a0e70c72c3b85.js.map",revision:"a361156c856d4f58f9fdd4f4685b4ef1"},{url:"/next/_next/static/chunks/855-bad530e5a6017dac.js",revision:"bad530e5a6017dac"},{url:"/next/_next/static/chunks/855-bad530e5a6017dac.js.map",revision:"cf78cd54485125852068026c474d87bd"},{url:"/next/_next/static/chunks/97cc2b9f-d1813b77a6d34ddf.js",revision:"d1813b77a6d34ddf"},{url:"/next/_next/static/chunks/97cc2b9f-d1813b77a6d34ddf.js.map",revision:"fed307c896f5e5a32ac98b7fa0451853"},{url:"/next/_next/static/chunks/e1533f8b-94b7dda262964aeb.js",revision:"94b7dda262964aeb"},{url:"/next/_next/static/chunks/e1533f8b-94b7dda262964aeb.js.map",revision:"4d74a43c2a784a232eaeb519c3d43dc1"},{url:"/next/_next/static/chunks/framework-0f3330e6efb75f4c.js",revision:"0f3330e6efb75f4c"},{url:"/next/_next/static/chunks/framework-0f3330e6efb75f4c.js.map",revision:"40683d0d7bb154b4d9601c9905b868b5"},{url:"/next/_next/static/chunks/main-3a0a9de2b38710bb.js",revision:"3a0a9de2b38710bb"},{url:"/next/_next/static/chunks/main-3a0a9de2b38710bb.js.map",revision:"5f241636b7ebb642ced154047c273869"},{url:"/next/_next/static/chunks/pages/_app-4fa74e3548f5d280.js",revision:"4fa74e3548f5d280"},{url:"/next/_next/static/chunks/pages/_app-4fa74e3548f5d280.js.map",revision:"b21aa0943739e0cfc5accb6c8f82bfa9"},{url:"/next/_next/static/chunks/pages/_error-ff3c94ec6a09d8c7.js",revision:"ff3c94ec6a09d8c7"},{url:"/next/_next/static/chunks/pages/_error-ff3c94ec6a09d8c7.js.map",revision:"14451aef01b1d16291874509ed655f9d"},{url:"/next/_next/static/chunks/pages/about-321407bc8f9353aa.js",revision:"321407bc8f9353aa"},{url:"/next/_next/static/chunks/pages/about-321407bc8f9353aa.js.map",revision:"e74c7aeb591c0a987f198765ac37f09e"},{url:"/next/_next/static/chunks/pages/admin/groups-01e03d06a7e19643.js",revision:"01e03d06a7e19643"},{url:"/next/_next/static/chunks/pages/admin/groups-01e03d06a7e19643.js.map",revision:"92d0f50a4a25a0cb94306ab66b950a4b"},{url:"/next/_next/static/chunks/pages/admin/links-9e96f3b543bbd057.js",revision:"9e96f3b543bbd057"},{url:"/next/_next/static/chunks/pages/admin/links-9e96f3b543bbd057.js.map",revision:"ec85d4cce2c74935a7e3d2465359d828"},{url:"/next/_next/static/chunks/pages/admin/users-888e1340f80c72b4.js",revision:"888e1340f80c72b4"},{url:"/next/_next/static/chunks/pages/admin/users-888e1340f80c72b4.js.map",revision:"dd8e1005d7e8893728791a1ffc70703f"},{url:"/next/_next/static/chunks/pages/auth/forgot-83cc3fbe7b37e002.js",revision:"83cc3fbe7b37e002"},{url:"/next/_next/static/chunks/pages/auth/forgot-83cc3fbe7b37e002.js.map",revision:"d19caa573c903e166617e7af7d62f5d1"},{url:"/next/_next/static/chunks/pages/auth/login-765927699ac8bb32.js",revision:"765927699ac8bb32"},{url:"/next/_next/static/chunks/pages/auth/login-765927699ac8bb32.js.map",revision:"e99b1b3e60d31e5d49a63d20edf0f347"},{url:"/next/_next/static/chunks/pages/auth/registration-8039d036f7421ba6.js",revision:"8039d036f7421ba6"},{url:"/next/_next/static/chunks/pages/auth/registration-8039d036f7421ba6.js.map",revision:"4ee28cded0541c660de5e103e23c6062"},{url:"/next/_next/static/chunks/pages/auth/verification-9615a77459f52ccd.js",revision:"9615a77459f52ccd"},{url:"/next/_next/static/chunks/pages/auth/verification-9615a77459f52ccd.js.map",revision:"7703c3dddb807076d040251f26194f71"},{url:"/next/_next/static/chunks/pages/contact-229c682351ac8d4f.js",revision:"229c682351ac8d4f"},{url:"/next/_next/static/chunks/pages/contact-229c682351ac8d4f.js.map",revision:"4b36037684009906f3f5eb06be927e54"},{url:"/next/_next/static/chunks/pages/faq-dbd9a7f580f54a96.js",revision:"dbd9a7f580f54a96"},{url:"/next/_next/static/chunks/pages/faq-dbd9a7f580f54a96.js.map",revision:"2d770cff529ccf7823db1b8123149536"},{url:"/next/_next/static/chunks/pages/index-ce8d0b63e8aa9a7d.js",revision:"ce8d0b63e8aa9a7d"},{url:"/next/_next/static/chunks/pages/index-ce8d0b63e8aa9a7d.js.map",revision:"63a3ce4538119e24d495c8a228615535"},{url:"/next/_next/static/chunks/pages/pricing-4711e46c291dcd01.js",revision:"4711e46c291dcd01"},{url:"/next/_next/static/chunks/pages/pricing-4711e46c291dcd01.js.map",revision:"54afb7127ed54e4c383d905c602318c1"},{url:"/next/_next/static/chunks/pages/privacy-e8b837037c727da1.js",revision:"e8b837037c727da1"},{url:"/next/_next/static/chunks/pages/privacy-e8b837037c727da1.js.map",revision:"152b7a908e026a684136533cf6cf28a8"},{url:"/next/_next/static/chunks/pages/user/addUrl-ab5676b37260ed73.js",revision:"ab5676b37260ed73"},{url:"/next/_next/static/chunks/pages/user/addUrl-ab5676b37260ed73.js.map",revision:"f38998e45ea4df335ba643c6385a935c"},{url:"/next/_next/static/chunks/pages/user/audit-cb4ec1a73fb61c69.js",revision:"cb4ec1a73fb61c69"},{url:"/next/_next/static/chunks/pages/user/audit-cb4ec1a73fb61c69.js.map",revision:"fb8a2b34b19f187915ccecb9ead69bd0"},{url:"/next/_next/static/chunks/pages/user/billing-ab5f9a7a1ad1b78c.js",revision:"ab5f9a7a1ad1b78c"},{url:"/next/_next/static/chunks/pages/user/billing-ab5f9a7a1ad1b78c.js.map",revision:"a5d5be4bb1a61c2eaf3819d502c54c7b"},{url:"/next/_next/static/chunks/pages/user/dashboard-da8387262a56a4f9.js",revision:"da8387262a56a4f9"},{url:"/next/_next/static/chunks/pages/user/dashboard-da8387262a56a4f9.js.map",revision:"e7a5acedd6531a6ab319c2425f294aad"},{url:"/next/_next/static/chunks/pages/user/integrations-407b08531abe0af0.js",revision:"407b08531abe0af0"},{url:"/next/_next/static/chunks/pages/user/integrations-407b08531abe0af0.js.map",revision:"e1e708e70ea9a1386dad75998dd1377e"},{url:"/next/_next/static/chunks/pages/user/links-78d66a5a23659e34.js",revision:"78d66a5a23659e34"},{url:"/next/_next/static/chunks/pages/user/links-78d66a5a23659e34.js.map",revision:"f392fafdc866fd307db8cf4dcc778f99"},{url:"/next/_next/static/chunks/pages/user/profile-d65357837aab2db0.js",revision:"d65357837aab2db0"},{url:"/next/_next/static/chunks/pages/user/profile-d65357837aab2db0.js.map",revision:"3896ac24d01edbcecc0b39f203373e48"},{url:"/next/_next/static/chunks/pages/user/reports-b5f5a7cc2b991048.js",revision:"b5f5a7cc2b991048"},{url:"/next/_next/static/chunks/pages/user/reports-b5f5a7cc2b991048.js.map",revision:"c0f2bcbf86173a1fbf957d046de8efa0"},{url:"/next/_next/static/chunks/polyfills-c67a75d1b6f99dc8.js",revision:"837c0df77fd5009c9e46d446188ecfd0"},{url:"/next/_next/static/chunks/webpack-7c2fda7efd65463d.js",revision:"7c2fda7efd65463d"},{url:"/next/_next/static/chunks/webpack-7c2fda7efd65463d.js.map",revision:"777efc861d211d9d61487173eeca3849"},{url:"/next/_next/static/css/1785585ae5de3930.css",revision:"1785585ae5de3930"},{url:"/next/_next/static/css/1785585ae5de3930.css.map",revision:"9753e7f933ace1d2c42c38d9cb4da3ef"},{url:"/next/_next/static/qCPBT5XCQsPOBHqV8_5N2/_buildManifest.js",revision:"affe41e3117f02513ff30f921972105e"},{url:"/next/_next/static/qCPBT5XCQsPOBHqV8_5N2/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/next/assets/images/undraw_back_in_the_day_knsh.svg",revision:"aebc6c499a138c3e107e65a208aec647"},{url:"/next/assets/images/undraw_co_workers_re_1i6i.svg",revision:"cb908c2f6d43c3d5bced6e0804dac2e9"},{url:"/next/assets/images/undraw_designer_re_5v95.svg",revision:"435c0b4cb909d0ceb63048a4e7ebc9f5"},{url:"/next/assets/images/undraw_welcome_cats_thqn.svg",revision:"ed0c3358facded075949f5e0ab20a080"},{url:"/next/assets/styles.css",revision:"02e913db9e4a418f8782e42d8195e42e"},{url:"/next/favicon.ico",revision:"c30c7d42707a47a3f4591831641e50dc"},{url:"/next/firebase-messaging-sw.js",revision:"47db0543b0c9d21608ee0cda826ce944"},{url:"/next/robots.txt",revision:"1b9a8511e7d3517e4b177176129da6f8"},{url:"/next/sitemap-0.xml",revision:"2ef8f99fa4fba2551f9effced9738793"},{url:"/next/sitemap.xml",revision:"1b97b3f5aaf3125af102e3090f13b360"}],{ignoreURLParametersMatching:[]}),e.cleanupOutdatedCaches(),e.registerRoute("/next",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({request:e,response:s,event:a,state:n})=>s&&"opaqueredirect"===s.type?new Response(s.body,{status:200,statusText:"OK",headers:s.headers}):s}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;const s=e.pathname;return!s.startsWith("/api/auth/")&&!!s.startsWith("/api/")}),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>{if(!(self.origin===e.origin))return!1;return!e.pathname.startsWith("/api/")}),new e.NetworkFirst({cacheName:"others",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:e})=>!(self.origin===e.origin)),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET")}));
//# sourceMappingURL=sw.js.map
