import { createApp } from "vue";
import { createPinia } from "pinia";
import App from "./App.vue";
import router from "./router";
import "./style.css";

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* import specific icons */
import { library } from '@fortawesome/fontawesome-svg-core';
import { faFacebook, faYoutube } from '@fortawesome/free-brands-svg-icons';
import { faEnvelope, faPhone, faMapMarkerAlt } from '@fortawesome/free-solid-svg-icons';

/* add icons to the library */
library.add(faFacebook, faYoutube, faEnvelope, faPhone, faMapMarkerAlt);

const app = createApp(App);

app.component('font-awesome-icon', FontAwesomeIcon);

app.use(createPinia());
app.use(router);

app.mount("#app");
