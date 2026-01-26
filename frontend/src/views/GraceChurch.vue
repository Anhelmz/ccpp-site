<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#101c33] via-[#152a47] to-[#1b3a5c] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 18% 25%,
                rgba(255, 255, 255, 0.2),
                transparent 45%
              ),
              radial-gradient(
                circle at 72% 15%,
                rgba(255, 255, 255, 0.15),
                transparent 35%
              ),
              radial-gradient(
                circle at 60% 80%,
                rgba(255, 255, 255, 0.18),
                transparent 30%
              );
          "
        ></div>
        <div
          class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"
        ></div>
      </div>
      <div
        class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-4"
      >
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
          {{ hero.title }}
        </h1>
        <p
          class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light"
        >
          {{ hero.subtitle }}
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="chip in hero.chips"
            :key="chip"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm"
          >
            {{ chip }}
          </span>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
            {{ intro.eyebrow }}
          </p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
            {{ intro.title }}
          </h2>
          <div
            class="space-y-4 text-left text-lg text-gray-700 leading-relaxed max-w-3xl mx-auto"
          >
            <p
              v-for="(paragraph, idx) in intro.body"
              :key="idx"
              v-html="paragraph"
            ></p>
          </div>
        </section>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <article
            v-for="card in cards"
            :key="card.id"
            class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4"
          >
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">
              {{ card.tag }}
            </p>
            <h3 class="text-2xl font-semibold text-gray-900">
              {{ card.title }}
            </h3>
            <p class="text-gray-600">
              {{ card.body }}
            </p>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            {{ cta.primary }}
          </router-link>
        </div>
      </div>
    </section>

    <!-- Gallery Section -->
    <section class="py-20 bg-gray-100">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-8">
        <div class="text-center space-y-3">
          <h2 class="text-4xl font-bold text-main">
            {{ galleryContent.title }}
          </h2>
          <p class="text-lg text-gray-700 max-w-2xl mx-auto">
            {{ galleryContent.description }}
          </p>
        </div>

        <div v-if="loading" class="text-center py-12">
          <div
            class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"
          ></div>
          <p class="text-gray-700">{{ galleryContent.loading }}</p>
        </div>
        <div
          v-else-if="galleryImages.length === 0"
          class="text-center text-gray-500 py-12"
        >
          {{ galleryContent.empty }}
        </div>
        <div v-else class="relative max-w-5xl mx-auto">
          <div
            ref="galleryScroll"
            class="flex gap-4 overflow-x-auto rounded-3xl shadow-2xl bg-gray-200 px-4 py-4 scroll-smooth"
          >
            <div
              v-for="(image, index) in galleryImages"
              :key="image.id || image.src"
              class="relative w-64 h-44 md:w-80 md:h-52 flex-shrink-0 rounded-2xl overflow-hidden bg-gray-200 cursor-pointer group"
              @click="openLightbox(index)"
            >
              <img
                :src="image.src"
                :alt="image.alt"
                class="w-full h-full object-cover group-hover:scale-110 transition-transform duration-500"
                @error="handleImageError"
              />
              <div
                class="absolute inset-0 bg-black/0 group-hover:bg-black/30 transition-all duration-300 flex items-center justify-center"
              >
                <svg
                  class="w-10 h-10 text-white opacity-0 group-hover:opacity-100 transition-opacity"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"
                  ></path>
                </svg>
              </div>
            </div>
          </div>
          <button
            class="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 hover:bg-white text-gray-700 rounded-full p-3 shadow-lg transition hidden md:flex"
            @click="scrollGallery(-1)"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 19l-7-7 7-7"
              />
            </svg>
          </button>
          <button
            class="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 hover:bg-white text-gray-700 rounded-full p-3 shadow-lg transition hidden md:flex"
            @click="scrollGallery(1)"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5l7 7-7 7"
              />
            </svg>
          </button>
        </div>
      </div>
    </section>

    <!-- Lightbox Modal -->
    <div
      v-if="lightboxOpen"
      class="fixed inset-0 bg-black/90 z-50 flex items-center justify-center p-4"
      @click="closeLightbox"
    >
      <button
        class="absolute top-4 right-4 text-white hover:text-gray-300 transition-colors z-10"
        @click="closeLightbox"
      >
        <svg
          class="w-8 h-8"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          ></path>
        </svg>
      </button>

      <button
        v-if="currentImageIndex > 0"
        class="absolute left-4 text-white hover:text-gray-300 transition-colors z-10"
        @click.stop="prevImage"
      >
        <svg
          class="w-12 h-12"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M15 19l-7-7 7-7"
          ></path>
        </svg>
      </button>

      <div class="max-w-5xl w-full" @click.stop>
        <img
          :src="galleryImages[currentImageIndex]?.src"
          :alt="galleryImages[currentImageIndex]?.alt"
          class="w-full h-auto max-h-[80vh] object-contain rounded-lg"
        />
      </div>

      <button
        v-if="currentImageIndex < galleryImages.length - 1"
        class="absolute right-4 text-white hover:text-gray-300 transition-colors z-10"
        @click.stop="nextImage"
      >
        <svg
          class="w-12 h-12"
          fill="none"
          stroke="currentColor"
          viewBox="0 0 24 24"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M9 5l7 7-7 7"
          ></path>
        </svg>
      </button>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import { galleryService } from "@/services/galleryService";

export default {
  name: "GraceChurch",
  setup() {
    const lightboxOpen = ref(false);
    const currentImageIndex = ref(0);
    const loading = ref(true);
    const error = ref("");
    const galleryImages = ref([]);
    const galleryScroll = ref(null);
    const galleryError = ref("Failed to load gallery images");

    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "Grace Church",
          subtitle:
            "House gatherings and Sunday services proclaiming the grace of Jesus across Phnom Penh.",
          chips: [
            "Word & Table",
            "House Gatherings",
            "Phnom Penh",
            "Grace Culture",
          ],
        },
        intro: {
          eyebrow: "Question",
          title: "What is Grace Church?",
          body: [
            'Grace Church gathers believers around <span class="font-semibold text-brand-orange">gospel-centered teaching</span>, worship, and community across Phnom Penh neighborhoods.',
            "We plant house churches, host Sunday services, and cultivate leaders who extend God’s grace to their friends, families, and co-workers.",
            "Every gathering keeps Scripture, communion, and Spirit-led prayer at the center so Jesus shapes our culture and mission.",
          ],
        },
        cards: [
          {
            id: "gather",
            tag: "Gather",
            title: "Word & worship",
            body: "Verse-by-verse teaching and gospel-saturated liturgy anchor every house gathering and Sunday service.",
          },
          {
            id: "community",
            tag: "Community",
            title: "Grace culture",
            body: "Hospitality, shared meals, and pastoral care create safe spaces to confess sin, celebrate redemption, and grow together.",
          },
          {
            id: "mission",
            tag: "Mission",
            title: "City-wide vision",
            body: "Training leaders and planting new communities ensures Phnom Penh hears the message of grace house by house.",
          },
        ],
        cta: { primary: "Visit Grace Church" },
        gallery: {
          title: "Gallery",
          description:
            "Scroll through images from Grace Church house gatherings, baptisms, and Sunday services.",
          loading: "Loading gallery...",
          empty: "No gallery images yet.",
          error: "Failed to load gallery images",
        },
      },
      kh: {
        hero: {
          title: "ព្រះវិហារព្រះគុណ",
          subtitle:
            "ក្រុមផ្ទះ និងការជួបជុំថ្ងៃអាទិត្យ ប្រកាសអំពីព្រះគុណរបស់ព្រះយេស៊ូវទូទាំងទីក្រុងភ្នំពេញ។",
          chips: [
            "ព្រះបន្ទូល និងព្រះជន្ម",
            "ក្រុមផ្ទះ",
            "ភ្នំពេញ",
            "វប្បធម៌ព្រះគុណ",
          ],
        },
        intro: {
          eyebrow: "សំណួរ",
          title: "តើព្រះវិហារព្រះគុណជាអ្វី?",
          body: [
            'ព្រះវិហារព្រះគុណ ប្រមូលអ្នកជឿជុំវិញ <span class="font-semibold text-brand-orange">ការបង្រៀនមជ្ឈមណ្ឌលព្រះសារ</span> ការថ្វាយបង្គំ និងសហគមន៍ តាមសង្កាត់ភ្នំពេញ។',
            "យើងដាំដុះក្រុមផ្ទះ រៀបចំការជួបជុំថ្ងៃអាទិត្យ និងបណ្តុះបណ្តាលអ្នកដឹកនាំ ដែលបន្តព្រះគុណរបស់ព្រះទៅកាន់មិត្ត ភារកិច្ច និងអ្នករួមការងារ។",
            "គ្រប់ការជួបជុំទាំងអស់ រក្សាព្រះគម្ពីរ ពិធីអាហារព្រះអម្ចាស់ និងការអធិស្ឋានដឹកនាំដោយព្រះវិញ្ញាណ នៅមជ្ឈមណ្ឌល ដើម្បីព្រះយេស៊ូវបង្កើតវប្បធម៌ និងភារកិច្ចរបស់យើង។",
          ],
        },
        cards: [
          {
            id: "gather",
            tag: "ជួបជុំ",
            title: "ព្រះបន្ទូល និងការថ្វាយបង្គំ",
            body: "ការបង្រៀនតាមជួរ និងពិធីអធិស្ឋានព្រះសារ បង្កប់គ្រប់ក្រុមផ្ទះ និងការជួបជុំថ្ងៃអាទិត្យ។",
          },
          {
            id: "community",
            tag: "សហគមន៍",
            title: "វប្បធម៌ព្រះគុណ",
            body: "ការស្វាគមន៍ អាហាររួម និងការថែរក្សាផ្លូវគង្វាល បង្កើតកន្លែងសុវត្ថិភាពសម្រាប់សារភាពបាប ប្រារព្ធរីករាយ និងលូតលាស់រួមគ្នា។",
          },
          {
            id: "mission",
            tag: "ភារកិច្ច",
            title: "វិស័យទាំងទីក្រុង",
            body: "ការបណ្តុះបណ្តាលអ្នកដឹកនាំ និងដាំដុះសហគមន៍ថ្មី ធានាថាទីក្រុងភ្នំពេញបានស្តាប់សារព្រះគុណពីផ្ទះទៅផ្ទះ។",
          },
        ],
        cta: { primary: "មកទស្សនាព្រះវិហារព្រះគុណ" },
        gallery: {
          title: "វិចិត្រសាល",
          description:
            "រុករករូបភាពពីក្រុមផ្ទះ ពិធីជ្រមុជទឹក និងការជួបជុំថ្ងៃអាទិត្យរបស់ព្រះវិហារព្រះគុណ។",
          loading: "កំពុងផ្ទុកវិចិត្រសាល...",
          empty: "មិនទាន់មានរូបវិចិត្រសាលទេ។",
          error: "មិនអាចផ្ទុកវិចិត្រសាលបានទេ",
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    const hero = computed(() => content.value.hero);
    const intro = computed(() => content.value.intro);
    const cards = computed(() => content.value.cards);
    const cta = computed(() => content.value.cta);
    const galleryContent = computed(() => content.value.gallery);

    const getImageUrl = (path) => {
      // Handle data URIs (data:image/...)
      if (path && path.startsWith("data:")) {
        return path;
      }
      // Handle HTTP/HTTPS URLs
      if (path && path.startsWith("http")) {
        return path;
      }
      // Handle relative paths
      return path && path.startsWith("/") ? path : `/${path || ''}`;
    };

    const fetchGalleryImages = async () => {
      loading.value = true;
      error.value = "";
      try {
        const response = await galleryService.getGalleries("grace-church");
        const galleries = response.galleries || [];
        galleryImages.value = galleries.map((gallery) => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `Grace Church Image ${gallery.id}`,
        }));
        loading.value = false;
      } catch (err) {
        console.error("Error loading gallery images:", err);
        error.value = galleryContent.value.error;
        loading.value = false;
      }
    };

    const handleImageError = (event) => {
      event.target.src =
        'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E';
    };

    const openLightbox = (index) => {
      currentImageIndex.value = index;
      lightboxOpen.value = true;
    };

    const closeLightbox = () => {
      lightboxOpen.value = false;
    };

    const nextImage = () => {
      if (currentImageIndex.value < galleryImages.value.length - 1) {
        currentImageIndex.value++;
      }
    };

    const prevImage = () => {
      if (currentImageIndex.value > 0) {
        currentImageIndex.value--;
      }
    };

    const scrollGallery = (direction) => {
      const el = galleryScroll.value;
      if (!el) return;
      const amount = el.clientWidth * 0.8 * direction;
      el.scrollBy({ left: amount, behavior: "smooth" });
    };

    onMounted(() => {
      fetchGalleryImages();
    });

    return {
      lightboxOpen,
      currentImageIndex,
      galleryImages,
      loading,
      error,
      handleImageError,
      galleryScroll,
      scrollGallery,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage,
      hero,
      intro,
      cards,
      cta,
      galleryContent,
    };
  },
};
</script>
