<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative min-h-[62vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-20"
    >
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="{
            backgroundImage: `linear-gradient(120deg, rgba(17, 39, 84, 0.25), rgba(17, 39, 84, 0.45)), url(${hero.image})`,
          }"
        ></div>
        <div
          class="absolute inset-0 bg-gradient-to-br from-black/70 via-black/55 to-black/65"
        ></div>
        <div
          class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-blue/35 via-brand-blue/15 to-transparent blur-[120px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-orange/30 via-brand-orange/12 to-transparent blur-[180px] opacity-65"
        ></div>
        <div
          class="absolute -top-24 left-1/2 w-[420px] h-[420px] -translate-x-1/2 bg-white/15 blur-[180px] opacity-60"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/80 via-black/60 to-transparent"
        ></div>
        <div
          class="absolute inset-0 opacity-15 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 20% 28%,
                rgba(255, 255, 255, 0.25),
                transparent 45%
              ),
              radial-gradient(
                circle at 80% 10%,
                rgba(255, 255, 255, 0.22),
                transparent 30%
              ),
              radial-gradient(
                circle at 62% 78%,
                rgba(255, 255, 255, 0.18),
                transparent 32%
              );
          "
        ></div>
      </div>
      <div
        class="relative z-10 w-full max-w-2xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-2"
      >
        <h1
          class="text-xl sm:text-2xl md:text-3xl font-semibold text-white leading-snug"
        >
          {{ hero.title }}
        </h1>
        <p class="text-xs sm:text-sm md:text-base text-white/80 font-light">
          {{ hero.subtitle }}
        </p>
      </div>
    </div>

    <!-- Story Section -->
    <section class="py-24 bg-white">
      <div
        class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid grid-cols-1 lg:grid-cols-2 gap-12"
      >
        <div>
          <p class="text-sm uppercase tracking-[0.3em] text-brand-blue mb-4">
            {{ story.eyebrow }}
          </p>
          <h2 class="text-4xl md:text-5xl font-bold text-brand-orange mb-6">
            {{ story.heading }}
          </h2>
          <p class="text-lg text-gray-700 leading-relaxed mb-4">
            {{ story.body[0] }}
          </p>
          <p class="text-lg text-gray-700 leading-relaxed">
            {{ story.body[1] }}
          </p>
        </div>
        <div
          class="relative rounded-3xl overflow-hidden border border-gray-100 shadow-2xl bg-gray-900/5"
        >
          <img
            :src="storyImage"
            :alt="story.imageAlt"
            class="w-full h-full object-cover"
          />
          <div
            class="absolute inset-0 bg-gradient-to-t from-black/45 via-transparent to-transparent"
          ></div>
        </div>
      </div>
    </section>

    <!-- Gallery Section -->
    <div class="relative py-24 text-white bg-brand-blue overflow-hidden">
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 bg-gradient-to-br from-brand-blue via-[#0a1f3e] to-[#153866]"
        ></div>
        <div
          class="absolute inset-y-[-25%] left-[-20%] w-2/3 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[160px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-30%] right-[-15%] w-1/2 bg-gradient-to-l from-white/15 via-transparent to-transparent blur-[140px] opacity-70"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-44 bg-gradient-to-t from-black/60 via-transparent to-transparent"
        ></div>
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 25% 20%,
                rgba(255, 255, 255, 0.25),
                transparent 45%
              ),
              radial-gradient(
                circle at 70% 10%,
                rgba(255, 255, 255, 0.22),
                transparent 35%
              ),
              radial-gradient(
                circle at 60% 80%,
                rgba(255, 255, 255, 0.18),
                transparent 30%
              );
          "
        ></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 space-y-16">
        <div class="text-center">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">
            {{ gallery.eyebrow }}
          </p>
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">
            {{ gallery.title }}
          </h2>
          <p class="text-lg text-white/80 max-w-2xl mx-auto font-light">
            {{ gallery.description }}
          </p>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
          <article
            v-for="highlight in highlightSections"
            :key="highlight.id"
            class="rounded-[32px] border border-white/15 bg-white/10 backdrop-blur-lg overflow-hidden flex flex-col md:flex-row"
          >
            <div class="md:w-1/2 relative">
              <img
                :src="highlight.image"
                :alt="highlight.title"
                class="w-full h-64 md:h-full object-cover"
              />
              <div
                class="absolute inset-0 bg-gradient-to-tr from-black/50 via-transparent to-transparent"
              ></div>
            </div>
            <div class="md:w-1/2 p-8 flex flex-col gap-4">
              <p class="text-xs uppercase tracking-[0.35em] text-white/70">
                {{ highlight.tag }}
              </p>
              <h3 class="text-2xl font-semibold leading-tight">
                {{ highlight.title }}
              </h3>
              <p class="text-white/80 text-base leading-relaxed">
                {{ highlight.description }}
              </p>
              <p class="text-white/60 text-sm leading-relaxed">
                {{ highlight.detail }}
              </p>
            </div>
          </article>
        </div>

        <div v-if="loading" class="text-center py-12">
          <div
            class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-white border-t-transparent mb-4"
          ></div>
          <p class="text-white/80">{{ gallery.loading }}</p>
        </div>
        <div v-else-if="error" class="text-center py-8 text-red-200">
          {{ error }}
        </div>
        <div
          v-else-if="galleryImages.length === 0"
          class="text-center text-white/60 py-12"
        >
          {{ gallery.empty }}
        </div>
        <div v-else class="relative max-w-5xl mx-auto">
          <div
            ref="galleryScroll"
            class="flex gap-4 overflow-x-auto rounded-3xl shadow-2xl bg-white/10 backdrop-blur-sm px-4 py-4 scroll-smooth"
          >
            <div
              v-for="(image, index) in galleryImages"
              :key="image.id || image.src"
              class="relative w-64 h-44 md:w-80 md:h-52 flex-shrink-0 rounded-2xl overflow-hidden bg-white/10 cursor-pointer group border border-white/20"
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
import { ref, onMounted, computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import { galleryService } from "@/services/galleryService";
import aboutHeroImage from "@/assets/background/bg-hero4.jpg";
import storyImage from "@/assets/background/background2.jpg";
import highlightWorship from "@/assets/gallery/547212625_1236675858494072_4425998334324483137_n.jpg";
import highlightOutreach from "@/assets/gallery/547491825_1236675745160750_5981939151572103357_n.jpg";
import highlightYouth from "@/assets/gallery/547541557_1236675618494096_195308602630672666_n.jpg";
import highlightPrayer from "@/assets/gallery/548044046_1236675671827424_5262319447121551375_n.jpg";

export default {
  name: "About",
  setup() {
    const lightboxOpen = ref(false);
    const currentImageIndex = ref(0);
    const loading = ref(true);
    const error = ref("");
    const galleryImages = ref([]);
    const galleryScroll = ref(null);
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "About Calvary Chapel Phnom Penh",
          subtitle:
            "A simple Calvary Chapel fellowship teaching the Bible verse by verse, welcoming Phnom Penh to know and follow Jesus together.",
          image: aboutHeroImage,
        },
        story: {
          eyebrow: "Our story",
          heading: "What is Calvary Chapel Phnom Penh?",
          body: [
            "Beginning with a handful of families studying Scripture together, the church has grown into a diverse community anchored in the whole counsel of God. We value expository preaching, Spirit-led worship, and discipleship for every generation.",
            "From Sunday gatherings to mid-week studies, outreach projects, and leadership development, our aim is to equip people to know Jesus, love one another, and serve Phnom Penh with humility and hope.",
          ],
          imageAlt: "Calvary Chapel Phnom Penh community",
        },
        gallery: {
          eyebrow: "Gallery",
          title: "Life around Calvary Chapel Phnom Penh",
          description:
            "Scenes from worship, discipleship, and everyday ministry in Phnom Penh.",
          loading: "Loading gallery...",
          empty: "No gallery images yet.",
          error: "Failed to load gallery images",
        },
        highlights: [
          {
            id: "worship",
            tag: "Gathered worship",
            title: "Sundays rooted in Scripture",
            description:
              "A bilingual family of believers worshiping, opening the Bible verse by verse, and sharing life together each week.",
            detail:
              "Captured during our 9:00 AM service inside the Phnom Penh sanctuary.",
            image: highlightWorship,
          },
          {
            id: "outreach",
            tag: "City outreach",
            title: "Serving Phnom Penh neighborhoods",
            description:
              "Teams distribute meals, pray with neighbors, and offer practical care throughout the city.",
            detail: "Moments from Serve Phnom Penh outreach mornings.",
            image: highlightOutreach,
          },
          {
            id: "youth",
            tag: "Next generation",
            title: "Youth discovering Jesus together",
            description:
              "Students gather midweek for teaching, discussion, and discipleship that equips them to follow Jesus boldly.",
            detail: "From our Friday youth ministry evenings.",
            image: highlightYouth,
          },
          {
            id: "prayer",
            tag: "Prayer & community",
            title: "House gatherings & prayer nights",
            description:
              "Circles of prayer, encouragement, and hospitality happening in homes across the city.",
            detail: "Snapshots from our monthly city prayer night.",
            image: highlightPrayer,
          },
        ],
      },
      kh: {
        hero: {
          title: "អំពី ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ",
          subtitle:
            "សហគមន៍ Calvary Chapel សាមញ្ញ បង្រៀនព្រះគម្ពីរតាមជួរ និងស្វាគមន៍ទីក្រុងភ្នំពេញឲ្យស្គាល់ និងតាមព្រះយេស៊ូវរួមគ្នា។",
          image: aboutHeroImage,
        },
        story: {
          eyebrow: "រឿងរ៉ាវរបស់យើង",
          heading: "តើ ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ ជាអ្វី?",
          body: [
            "ចាប់ផ្តើមពីគ្រួសារតិចតួចដែលសិក្សាព្រះគម្ពីររួមគ្នា គ្រឹស្តសាសនាចក្រ បានលូតលាស់ក្លាយជាសហគមន៍ចម្រុះ ដែលជាប់នឹងសេចក្តីប្រកាសព្រះវចនៈទាំងមូល។ យើងពឹងផ្អែកលើការប្រាប់អត្ថន័យព្រះគម្ពីរ ការថ្វាយបង្គំដឹកនាំដោយព្រះវិញ្ញាណ និងការបណ្តុះបណ្តាលសិស្សសម្រាប់គ្រួសារទាំងអស់។",
            "ចាប់ពីការជួបជុំថ្ងៃអាទិត្យ ដល់សិក្សាកណ្ដាលសប្ដាហ៍ គម្រោងចេញផ្សាយជួយ និងការអភិវឌ្ឍភាពជាអ្នកដឹកនាំ គោលបំណងរបស់យើងគឺបំប៉នមនុស្សឲ្យស្គាល់ព្រះយេស៊ូវ ស្រឡាញ់គ្នា និងបម្រើទីក្រុងភ្នំពេញដោយភាពទាប និងសេចក្តីសង្ឃឹម។",
          ],
          imageAlt: "សហគមន៍ ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ",
        },
        gallery: {
          eyebrow: "វិចិត្រសាល",
          title: "ជីវិតជុំវិញ ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ",
          description:
            "រូបភាពពីការថ្វាយបង្គំ ការបណ្តុះសិស្ស និងសេវាកម្មប្រចាំថ្ងៃនៅភ្នំពេញ។",
          loading: "កំពុងផ្ទុកវិចិត្រសាល...",
          empty: "មិនទាន់មានរូបវិចិត្រសាលទេ។",
          error: "មិនអាចផ្ទុកវិចិត្រសាលបានទេ",
        },
        highlights: [
          {
            id: "worship",
            tag: "ការថ្វាយបង្គំរួមគ្នា",
            title: "ថ្ងៃអាទិត្យជាប់នឹងព្រះវចនៈ",
            description:
              "គ្រួសារជឿពីភាសាទាំងពីរថ្វាយបង្គំ បើកព្រះគម្ពីរតាមជួរ និងចែករំលែកជីវិតជាមួយគ្នារៀងរាល់សប្ដាហ៍។",
            detail: "ថតបានក្នុងពេលបំពេញបេសកកម្មម៉ោង 9 ព្រឹក នៅសាលធំភ្នំពេញ។",
            image: highlightWorship,
          },
          {
            id: "outreach",
            tag: "ការចេញផ្សាយទីក្រុង",
            title: "បម្រើសង្កាត់ក្នុងទីក្រុងភ្នំពេញ",
            description:
              "ក្រុមចែកចាយអាហារ អធិស្ឋានជាមួយអ្នកជិតខាង និងផ្តល់ការថែទាំជាក់ស្តែងទូទាំងទីក្រុង។",
            detail: "ពេលវេលាពីព្រឹត្តិការណ៍ Serve Phnom Penh ពេលព្រឹក។",
            image: highlightOutreach,
          },
          {
            id: "youth",
            tag: "ជំនាន់ក្រោយ",
            title: "យុវជនស្វែងយល់ពីព្រះយេស៊ូវរួមគ្នា",
            description:
              "និស្សិតជួបជុំកណ្ដាលសប្ដាហ៍ សម្រាប់ការបង្រៀន ការពិភាក្សា និងការបណ្តុះបណ្តាលសិស្ស ដើម្បីដើរតាមព្រះយេស៊ូវដោយក្លាហាន។",
            detail: "ពីយប់សេវាយុវជនថ្ងៃសុក្រ។",
            image: highlightYouth,
          },
          {
            id: "prayer",
            tag: "ការអធិស្ឋាន និងសហគមន៍",
            title: "ក្រុមផ្ទះ និងយប់អធិស្ឋាន",
            description:
              "រង្វង់អធិស្ឋាន ការលើកទឹកចិត្ត និងការស្វាគមន៍នៅគេហដ្ឋានជុំវិញទីក្រុង។",
            detail: "រូបពីយប់អធិស្ឋានទីក្រុងប្រចាំខែ។",
            image: highlightPrayer,
          },
        ],
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    const hero = computed(() => content.value.hero);
    const story = computed(() => content.value.story);
    const gallery = computed(() => content.value.gallery);
    const highlightSections = computed(() => content.value.highlights);

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
        const response = await galleryService.getGalleries("about");
        const galleries = response.galleries || [];
        galleryImages.value = galleries.map((gallery) => ({
          id: gallery.id,
          src: getImageUrl(gallery.path),
          alt: gallery.filename || `About Image ${gallery.id}`,
        }));
        loading.value = false;
      } catch (err) {
        console.error("Error loading gallery images:", err);
        error.value = gallery.value.error;
        loading.value = false;
      }
    };

    const handleImageError = (event) => {
      // Hide the image when it fails to load instead of showing placeholder
      event.target.style.display = 'none';
      // Also hide the parent container to remove the blank space
      if (event.target.parentElement) {
        event.target.parentElement.style.display = 'none';
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
      story,
      storyImage,
      gallery,
      highlightSections,
    };
  },
};
</script>
