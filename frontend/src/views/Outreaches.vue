<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#0a1f33] via-[#102a47] to-[#153a5d] text-white"
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

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="araksatImage"
                alt="Araksat Outreach"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div
                class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"
              ></div>
              <span
                class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900"
              >
                {{ cards[0].tag }}
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">
                {{ cards[0].title }}
              </h3>
              <p class="text-gray-600 flex-1">
                {{ cards[0].body }}
              </p>
            </div>
          </article>

          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="steunmeancheyImage"
                alt="Steunmeanchey Outreach"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div
                class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"
              ></div>
              <span
                class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900"
              >
                {{ cards[1].tag }}
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">
                {{ cards[1].title }}
              </h3>
              <p class="text-gray-600 flex-1">
                {{ cards[1].body }}
              </p>
            </div>
          </article>

          <article
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="preyVengImage"
                alt="Prey Veng Outreach"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div
                class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"
              ></div>
              <span
                class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900"
              >
                {{ cards[2].tag }}
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">
                {{ cards[2].title }}
              </h3>
              <p class="text-gray-600 flex-1">
                {{ cards[2].body }}
              </p>
            </div>
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
    <div class="py-20 bg-gray-200">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <h2 class="text-4xl font-bold text-center text-main mb-4">Gallery</h2>
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Glimpses of our outreach programs and community service
        </p>

        <div v-if="loading" class="text-center py-12">
          <div
            class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-main border-t-transparent mb-4"
          ></div>
          <p class="text-gray-700">Loading gallery...</p>
        </div>

        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-600">{{ error }}</p>
        </div>

        <div v-else>
          <div
            class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-8"
          >
            <div
              v-for="(image, index) in paginatedImages"
              :key="image.id || image.src"
              class="relative overflow-hidden rounded-lg shadow-lg hover:shadow-2xl transition-all duration-300 cursor-pointer group bg-gray-200"
              @click="openLightbox(getFullImageIndex(index))"
            >
              <img
                :src="image.src"
                :alt="image.alt"
                class="w-full h-64 object-cover group-hover:scale-110 transition-transform duration-500"
                @error="handleImageError"
              />
              <div
                class="absolute inset-0 bg-black/0 group-hover:bg-black/30 transition-all duration-300 flex items-center justify-center"
              >
                <svg
                  class="w-12 h-12 text-white opacity-0 group-hover:opacity-100 transition-opacity"
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

          <!-- Pagination Controls -->
          <div
            v-if="totalPages > 1"
            class="flex items-center justify-center gap-2 mt-8"
          >
            <button
              :disabled="currentPage === 1"
              class="px-4 py-2 bg-main text-white rounded-lg hover:bg-secondary transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              @click="goToPage(currentPage - 1)"
            >
              Previous
            </button>

            <div class="flex gap-2">
              <button
                v-for="page in totalPages"
                :key="page"
                :class="[
                  'px-4 py-2 rounded-lg transition-colors',
                  currentPage === page
                    ? 'bg-brand-orange text-white font-semibold'
                    : 'bg-gray-200 text-gray-700 hover:bg-gray-300',
                ]"
                @click="goToPage(page)"
              >
                {{ page }}
              </button>
            </div>

            <button
              :disabled="currentPage === totalPages"
              class="px-4 py-2 bg-main text-white rounded-lg hover:bg-secondary transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              @click="goToPage(currentPage + 1)"
            >
              Next
            </button>
          </div>
        </div>
      </div>
    </div>

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
import araksatImage from "@/assets/gallery/548625714_1236675678494090_1776087346866151382_n.jpg";
import steunmeancheyImage from "@/assets/gallery/547653470_1236675908494067_1828648105188893790_n.jpg";
import preyVengImage from "@/assets/gallery/548222744_1236675851827406_3818316444684279685_n.jpg";

export default {
  name: "Outreaches",
  setup() {
    const lightboxOpen = ref(false);
    const currentImageIndex = ref(0);
    const currentPage = ref(1);
    const itemsPerPage = 6;
    const loading = ref(true);
    const error = ref("");
    const galleryImages = ref([]);
    const galleryError = ref("Failed to load gallery images");

    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "Outreaches",
          subtitle:
            "Serving Cambodia with tangible compassion, prayer, and the hope of Jesus.",
          chips: ["MercyTeams", "Meals & Care", "Prayer", "Gospel Moments"],
        },
        intro: {
          eyebrow: "Question",
          title: "What are our outreaches?",
          body: [
            'Outreaches carry <span class="font-semibold text-brand-orange">Calvary Chapel Phnom Penh</span> into neighborhoods, schools, and villages with meals, medical care, prayer, and discipleship.',
            'We partner with local leaders, listen to community needs, and respond with <span class="font-semibold text-brand-orange">compassionate action</span> that points every person to Jesus.',
            'Whether we’re serving families, youth, or vulnerable neighbors, our teams keep the <span class="font-semibold text-brand-orange">Gospel and presence of the Holy Spirit</span> at the center.',
          ],
        },
        cards: [
          {
            id: "araksat",
            tag: "Araksat",
            title: "Village gatherings",
            body: "Meals, house visits, and Bible studies anchor this long-term relationship with families north of the city.",
          },
          {
            id: "steung-meanchey",
            tag: "Steung Meanchey",
            title: "Urban compassion",
            body: "Care packages, medical clinics, and youth ministry support families living around the landfill district.",
          },
          {
            id: "prey-veng",
            tag: "Prey Veng",
            title: "Regional partnership",
            body: "Monthly teams encourage rural churches with leadership training, prayer, and evangelistic outreaches.",
          },
        ],
        cta: { primary: "Serve with Us" },
        galleryError: "Failed to load gallery images",
      },
      kh: {
        hero: {
          title: "ការចេញផ្សាយ",
          subtitle:
            "បម្រើកម្ពុជា​ដោយមេត្តាករុណាជាក់ស្តែង ការអធិស្ឋាន និងសេចក្តីសង្ឃឹមក្នុងព្រះយេស៊ូវ។",
          chips: [
            "ក្រុមមេត្តាករុណា",
            "អាហារ និងការថែទាំ",
            "ការអធិស្ឋាន",
            "ខណៈពេលសារព្រះ",
          ],
        },
        intro: {
          eyebrow: "សំណួរ",
          title: "ការចេញផ្សាយរបស់យើងមានអ្វីខ្លះ?",
          body: [
            'ការចេញផ្សាយនាំ <span class="font-semibold text-brand-orange">ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ</span> ចូលទៅសង្កាត់ សាលារៀន និងភូមិ ដោយនាំអាហារ ការថែទាំសុខភាព ការអធិស្ឋាន និងការបណ្តុះសិស្ស។',
            'យើងសហការជាមួយអ្នកដឹកនាំក្នុងស្រុក ស្តាប់តម្រូវការសហគមន៍ និងឆ្លើយតបដោយ <span class="font-semibold text-brand-orange">សកម្មភាពមេត្តាករុណា</span> ដែលបង្ហាញមនុស្សគ្រប់រូបទៅរកព្រះយេស៊ូវ។',
            'មិនថាកំពុងបម្រើគ្រួសារ យុវជន ឬអ្នកងាយរងគ្រោះ ក្រុមរបស់យើងរក្សា <span class="font-semibold text-brand-orange">សារព្រះ និងវត្តមានព្រះវិញ្ញាណបរិសុទ្ធ</span> ជាមជ្ឈមណ្ឌលជានិច្ច។',
          ],
        },
        cards: [
          {
            id: "araksat",
            tag: "អារាក់សាត់",
            title: "ការជួបជុំព្រះវិហារ​ភូមិ",
            body: "អាហារ ការទៅលេងផ្ទះ និងការសិក្សាព្រះគម្ពីរ ជាគ្រឹះសម្រាប់ទំនាក់ទំនងរយៈពេលវែងជាមួយគ្រួសារទិសខាងជើងទីក្រុង។",
          },
          {
            id: "steung-meanchey",
            tag: "ស្ទឹងមានជ័យ",
            title: "មេត្តាករុណាក្រុង",
            body: "កញ្ចប់ជំនួយ គ្លីនិកសុខភាព និងសេវាយុវជន គាំទ្រគ្រួសារនៅជុំវិញតំបន់ល្បាក់សំរាម។",
          },
          {
            id: "prey-veng",
            tag: "ព្រៃវែង",
            title: "ដៃគូតំបន់",
            body: "ក្រុមប្រចាំខែ បំផុសគ្រឹស្តសាសនាចក្រ​ជនបទតាមការបណ្តុះបណ្តាលភាពជាអ្នកដឹកនាំ ការអធិស្ឋាន និងការចេញផ្សាយសារ​សង្គ្រោះ។",
          },
        ],
        cta: { primary: "ចូលរួមបម្រើជាមួយយើង" },
        galleryError: "មិនអាចផ្ទុកវិចិត្រសាលបានទេ",
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    galleryError.value = content.value.galleryError;
    const hero = computed(() => content.value.hero);
    const intro = computed(() => content.value.intro);
    const cards = computed(() => content.value.cards);
    const cta = computed(() => content.value.cta);

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
        const response = await galleryService.getGalleries("outreaches");
        const galleries = response.galleries || [];
        galleryImages.value = galleries.map((gallery) => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `Outreach Image ${gallery.id}`,
        }));
        loading.value = false;
      } catch (err) {
        console.error("Error loading gallery images:", err);
        error.value = galleryError.value;
        loading.value = false;
      }
    };

    const handleImageError = (event) => {
      event.target.src =
        'data:image/svg+xml,%3Csvg xmlns="http://www.w3.org/2000/svg" width="200" height="200"%3E%3Crect fill="%23ddd" width="200" height="200"/%3E%3Ctext fill="%23999" font-family="sans-serif" font-size="14" dy="10.5" font-weight="bold" x="50%25" y="50%25" text-anchor="middle"%3EImage not found%3C/text%3E%3C/svg%3E';
    };

    const totalPages = computed(() => {
      return Math.ceil(galleryImages.value.length / itemsPerPage);
    });

    const paginatedImages = computed(() => {
      const start = (currentPage.value - 1) * itemsPerPage;
      const end = start + itemsPerPage;
      return galleryImages.value.slice(start, end);
    });

    const getFullImageIndex = (paginatedIndex) => {
      return (currentPage.value - 1) * itemsPerPage + paginatedIndex;
    };

    const goToPage = (page) => {
      if (page >= 1 && page <= totalPages.value) {
        currentPage.value = page;
        window.scrollTo({ top: 0, behavior: "smooth" });
      }
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

    onMounted(() => {
      fetchGalleryImages();
    });

    return {
      lightboxOpen,
      currentImageIndex,
      galleryImages,
      loading,
      error,
      openLightbox,
      closeLightbox,
      nextImage,
      prevImage,
      currentPage,
      itemsPerPage,
      totalPages,
      paginatedImages,
      getFullImageIndex,
      goToPage,
      handleImageError,
      araksatImage,
      steunmeancheyImage,
      preyVengImage,
      hero,
      intro,
      cards,
      cta,
      galleryError,
    };
  },
};
</script>
