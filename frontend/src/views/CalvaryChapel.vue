<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[70vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-20 bg-gradient-to-br from-[#081229] via-[#0f233f] to-[#173452] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-30 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 20% 20%,
                rgba(255, 255, 255, 0.25),
                transparent 40%
              ),
              radial-gradient(
                circle at 80% 15%,
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
          class="absolute inset-y-[-30%] left-[-15%] w-2/3 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[180px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-20%] right-[-10%] w-1/2 bg-gradient-to-l from-brand-blue/35 via-transparent to-transparent blur-[150px] opacity-70"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/40 via-transparent to-transparent"
        ></div>
      </div>
      <div
        class="relative z-10 w-full max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6"
      >
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
          {{ hero.title }}
        </h1>
        <p
          class="text-lg md:text-xl text-white/80 max-w-3xl mx-auto font-light"
        >
          {{ hero.subtitle }}
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="point in heroPoints"
            :key="point"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.3em] backdrop-blur-sm"
          >
            {{ point }}
          </span>
        </div>
      </div>
    </section>

    <!-- About Section -->
    <section class="py-20 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 space-y-6 text-center">
        <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
          {{ about.eyebrow }}
        </p>
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
          {{ about.title }}
        </h2>
        <div class="space-y-4 text-left text-lg text-gray-600 leading-relaxed">
          <p v-for="(paragraph, idx) in about.body" :key="idx">
            {{ paragraph }}
          </p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";

export default {
  name: "CalvaryChapel",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "Calvary Chapel",
          subtitle:
            "A global fellowship of churches rooted in expositional teaching, the work of the Holy Spirit, and a grace-filled culture that welcomes every generation.",
        },
        heroPoints: [
          "Verse-by-verse",
          "Spirit-led",
          "Global family",
          "Grace & truth",
        ],
        about: {
          eyebrow: "Question",
          title: "What is Calvary Chapel?",
          body: [
            "Calvary Chapel is an international association of charismatic evangelical churches with roots in the 1960s Jesus Movement. Beginning in Costa Mesa, California, the fellowship embraced verse-by-verse Bible teaching, simple worship, and a welcoming culture that drew thousands of new believers.",
            "The churches identify as neither a denomination nor strictly Pentecostal. They hold to historic evangelical doctrine, practice the gifts of the Holy Spirit in an orderly way, and maintain a strong expectation of Christ’s return.",
            "Today the Calvary Chapel family includes radio ministries, Bible colleges, and church plants around the world—all focused on teaching Scripture, depending on the Spirit, and serving communities with the hope of Jesus.",
          ],
        },
      },
      kh: {
        hero: {
          title: "កាល់វ៉ារីឆាបផល",
          subtitle:
            "សហគមន៍សកលនៃសាសនាចក្រ ដែលជាប់នឹងការបង្រៀនតាមជួរ ការធ្វើការរបស់ព្រះវិញ្ញាណបរិសុទ្ធ និងវប្បធម៌ពោរពេញដោយព្រះគុណដែលស្វាគមន៍គ្រប់ជំនាន់។",
        },
        heroPoints: [
          "បង្រៀនតាមជួរ",
          "ដឹកនាំដោយព្រះវិញ្ញាណ",
          "គ្រួសារសកល",
          "ព្រះគុណ និងសេចក្តីពិត",
        ],
        about: {
          eyebrow: "សំណួរ",
          title: "តើ កាល់វ៉ារីឆាបផល ជាអ្វី?",
          body: [
            "កាល់វ៉ារីឆាបផល ជាសមាគមអន្តរជាតិនៃសាសនាចក្រអេវ៉ានជីលីកដែលមានអំណោយអារម្មណ៍ មានដើមកំណើតពីចលនាព្រះយេស៊ូវក្នុងទសវត្សរ៍ 1960។ ចាប់ផ្តើមនៅ Costa Mesa រដ្ឋកាលីហ្វോർនៀ សហគមន៍បានអើពើការបង្រៀនព្រះគម្ពីរតាមជួរ ការថ្វាយបង្គំសាមញ្ញ និងវប្បធម៌ស្វាគមន៍ដែលទាក់ទាញអ្នកជឿថ្មីរាប់ពាន់នាក់។",
            "សាសនាចក្រទាំងនេះមិនកំណត់ខ្លួនជានិកាយ ឬពាក់ព័ន្ធពូរវប្បធម៌ទេ។ ពួកគេចាំខ្ជាប់នូវសេចក្តីបង្រៀនអេវ៉ានជីលីកបុរាណ អនុវត្តអំណោយព្រះវិញ្ញាណបរិសុទ្ធដោយសណ្ឋិត និងរក្សាការរំពឹងជឿលើការត្រឡប់មកវិញរបស់ព្រះគ្រីស្ទ។",
            "សព្វថ្ងៃ គ្រួសារ កាល់វ៉ារីឆាបផល រួមមានវិទ្យុ សាលាព្រះគម្ពីរ និងការដាំដុះសាសនាចក្រ ជុំវិញពិភពលោក—ផ្តោតលើការបង្រៀនព្រះគម្ពីរ ពឹងព្រះវិញ្ញាណ និងបម្រើសហគមន៍ដោយសេចក្តីសង្ឃឹមនៅក្នុងព្រះយេស៊ូវ។",
          ],
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );

    return {
      hero: computed(() => content.value.hero),
      heroPoints: computed(() => content.value.heroPoints),
      about: computed(() => content.value.about),
    };
  },
};
</script>
