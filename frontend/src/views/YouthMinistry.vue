<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#141a38] via-[#1c2550] to-[#253368] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 15% 25%,
                rgba(255, 255, 255, 0.2),
                transparent 45%
              ),
              radial-gradient(
                circle at 70% 15%,
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
          class="absolute inset-y-[-25%] left-[-8%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-65"
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
        <div v-else-if="error" class="text-center py-8 text-red-600">
          {{ error }}
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
                class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-110"
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
  name: "YouthMinistry",
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
          title: "Youth Ministry",
          subtitle:
            "Students discovering who they are in Jesus through Scripture, community, and mission.",
          chips: [
            "Fridays 6:30PM",
            "Ages 12-18",
            "Scripture & Worship",
            "City Serve",
          ],
        },
        intro: {
          eyebrow: "Question",
          title: "What is Youth Ministry?",
          body: [
            'Friday nights are designed for <span class="font-semibold text-brand-orange">teens to encounter Jesus</span>, ask honest questions, and build friendships that last beyond the gathering.',
            'We teach the Bible directly, invite students into <span class="font-semibold text-brand-orange">Spirit-led worship</span>, and create space for conversations that connect faith to Phnom Penh life.',
            "Serving together—from city outreaches to helping on Sundays—forms students into disciples who love Jesus and their neighbors.",
          ],
        },
        cards: [
          {
            id: "teaching",
            tag: "Teaching",
            title: "Bible-centered nights",
            body: "Verse-by-verse teaching, small-group discussion, and mentoring conversations help students apply Scripture.",
          },
          {
            id: "community",
            tag: "Community",
            title: "Shared life & worship",
            body: "Student-led worship, hangouts, and retreats build a healthy community where every teen belongs.",
          },
          {
            id: "mission",
            tag: "Mission",
            title: "Serve Phnom Penh",
            body: "Monthly projects, prayer walks, and partnering with outreaches teach students to live on mission daily.",
          },
        ],
        cta: { primary: "Connect a student" },
        gallery: {
          title: "Gallery",
          description:
            "Scroll through Youth Ministry moments from gatherings, retreats, and serve days.",
          loading: "Loading gallery...",
          empty: "No gallery images yet.",
          error: "Failed to load gallery images",
        },
      },
      kh: {
        hero: {
          title: "សេវាយុវជន",
          subtitle:
            "យុវជនស្វែងរកអត្តសញ្ញាណក្នុងព្រះយេស៊ូវ តាមរយៈព្រះវចនៈ សហគមន៍ និងភារកិច្ច។",
          chips: [
            "ថ្ងៃសុក្រ 6:30 ល្ងាច",
            "អាយុ 12-18",
            "ព្រះវចនៈ និងការថ្វាយបង្គំ",
            "បម្រើទីក្រុង",
          ],
        },
        intro: {
          eyebrow: "សំណួរ",
          title: "តើសេវាយុវជនគឺអ្វី?",
          body: [
            'យប់ថ្ងៃសុក្រ ត្រូវបានរៀបចំសម្រាប់ <span class="font-semibold text-brand-orange">យុវជនឱ្យជួបព្រះយេស៊ូវ</span> សួរសំណួរពិតប្រាកដ និងស្ថាបនាមិត្តភាពដែលតែងតែបន្តក្រៅការជួបជុំ។',
            'យើងបង្រៀនព្រះគម្ពីរត្រង់ៗ អញ្ជើញយុវជនចូលរៀងរាល់ <span class="font-semibold text-brand-orange">ការថ្វាយបង្គំដឹកនាំដោយព្រះវិញ្ញាណ</span> និងបង្កើតកន្លែងសន្ទនាដែលភ្ជាប់ជំនឿនឹងជីវិតនៅទីក្រុងភ្នំពេញ។',
            "ការបម្រើរួមគ្នា ចាប់ពីការចេញផ្សាយក្នុងទីក្រុង ដល់ជួយថ្ងៃអាទិត្យ បង្កើតយុវជនឱ្យក្លាយជាសិស្សដែលស្រឡាញ់ព្រះយេស៊ូវ និងអ្នកជិតខាង។",
          ],
        },
        cards: [
          {
            id: "teaching",
            tag: "ការបង្រៀន",
            title: "យប់ប្រកបដោយព្រះវចនៈ",
            body: "ការបង្រៀនតាមជួរ ក្រុមតូច និងការជួបពិភាក្សា ជួយយុវជនអនុវត្តព្រះគម្ពីរ។",
          },
          {
            id: "community",
            tag: "សហគមន៍",
            title: "ជីវិត និងការថ្វាយបង្គំរួម",
            body: "ការថ្វាយបង្គំដឹកនាំដោយយុវជន ការស្នាក់ស្នាល និងជំរះ បង្កើតសហគមន៍មានសុខភាពដែលយុវជនគ្រប់រូបមានទីកន្លែង។",
          },
          {
            id: "mission",
            tag: "ភារកិច្ច",
            title: "បម្រើទីក្រុងភ្នំពេញ",
            body: "គម្រោងប្រចាំខែ ដំណើរការអធិស្ឋាន និងសហការជាមួយការចេញផ្សាយ បង្រៀនយុវជនឱ្យរស់នៅជាអ្នកបម្រើរៀងរាល់ថ្ងៃ។",
          },
        ],
        cta: { primary: "ភ្ជាប់យុវជនម្នាក់" },
        gallery: {
          title: "វិចិត្រសាល",
          description: "រុករករូបភាពពីការជួបជុំយុវជន ជំនួបជំរះ និងថ្ងៃបម្រើ។",
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
      if (path.startsWith("http")) return path;
      return path.startsWith("/") ? path : `/${path}`;
    };

    const fetchGalleryImages = async () => {
      loading.value = true;
      error.value = "";
      try {
        const response = await galleryService.getGalleries("youth");
        const galleries = response.galleries || [];
        galleryImages.value = galleries.map((gallery) => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `Youth Ministry Image ${gallery.id}`,
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
