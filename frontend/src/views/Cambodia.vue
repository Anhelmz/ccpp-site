<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[70vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-20 bg-gradient-to-br from-[#0b1c34] via-[#142b4a] to-[#1d3c5f] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-30 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 25% 20%,
                rgba(255, 255, 255, 0.18),
                transparent 45%
              ),
              radial-gradient(
                circle at 75% 15%,
                rgba(255, 255, 255, 0.12),
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
          class="absolute inset-y-[-25%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[180px] opacity-65"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/50 via-black/30 to-transparent"
        ></div>
      </div>
      <div
        class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6"
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
          <p
            v-for="(paragraph, idx) in about.body"
            :key="idx"
            v-html="paragraph"
          ></p>
        </div>
      </div>
    </section>
    <section class="pb-16 bg-white">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <a
          href="https://en.wikipedia.org/wiki/Cambodia"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
        >
          {{ about.cta }}
        </a>
      </div>
    </section>
  </div>
</template>

<script>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";

export default {
  name: "Cambodia",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "Cambodia",
          subtitle:
            "A vibrant nation of resilience, heritage, and Gospel opportunity in the heart of Southeast Asia.",
          chips: ["Phnom Penh", "Khmer Heritage", "Tonlé Sap", "Gospel Future"],
        },
        about: {
          eyebrow: "Question",
          title: "What is Cambodia?",
          body: [
            '<span class="font-semibold text-brand-orange">Cambodia</span>, officially the <span class="font-semibold text-brand-orange">Kingdom of Cambodia</span>, sits at the crossroads of <span class="font-semibold text-brand-orange">Southeast Asia</span> bordered by Thailand, Laos, Vietnam, and the Gulf of Thailand. The <span class="font-semibold text-brand-orange">Mekong River</span>, <span class="font-semibold text-brand-orange">Tonlé Sap Lake</span>, and fertile plains shape daily life for the 17+ million people who call the country home.',
            '<span class="font-semibold text-brand-orange">Phnom Penh</span>, the capital, pulses at the confluence of the Mekong and Tonlé Sap rivers. From here, <span class="font-semibold text-brand-orange">Calvary Chapel Phnom Penh</span> shares the hope of Jesus within a city marked by markets, cafes, universities, and local neighborhoods that represent every layer of Cambodian society.',
            'The legacy of the <span class="font-semibold text-brand-orange">Khmer Empire</span>—highlighted by <span class="font-semibold text-brand-orange">Angkor Wat</span>—still inspires national pride. Yet recent history includes deep trauma from the <span class="font-semibold text-brand-orange">Khmer Rouge</span> regime. Today’s Cambodia is rebuilding through education, entrepreneurship, and a new generation eager for truth, restoration, and purpose.',
            '<span class="font-semibold text-brand-orange">Ministry in Cambodia</span> involves listening well, honoring the culture, investing in relationships, and presenting the <span class="font-semibold text-brand-orange">Gospel</span> faithfully in Khmer. From house gatherings to city outreaches, we seek to equip believers and serve Phnom Penh with humility and hope.',
          ],
          cta: "Learn More",
        },
      },
      kh: {
        hero: {
          title: "កម្ពុជា",
          subtitle:
            "ប្រទេសពោរពេញដោយភាពធន់ មរតកវប្បធម៌ និងឱកាសសម្រាប់សារព្រះ នៅចំណុចកណ្ដាលអាស៊ីអាគ្នេយ៍។",
          chips: ["ភ្នំពេញ", "មរតកខ្មែរ", "ទន្លេសាប", "អនាគតសារព្រះ"],
        },
        about: {
          eyebrow: "សំណួរ",
          title: "តើកម្ពុជាជាអ្វី?",
          body: [
            '<span class="font-semibold text-brand-orange">កម្ពុជា</span> ឬ <span class="font-semibold text-brand-orange">ព្រះរាជាណាចក្រកម្ពុជា</span> ស្ថិតនៅចំណុចឆ្លងកាត់នៃ <span class="font-semibold text-brand-orange">អាស៊ីអាគ្នេយ៍</span> ជិតថៃ ឡាវ វៀតណាម និងឈូងសមុទ្រថៃ។ <span class="font-semibold text-brand-orange">ទន្លេមេគង្គ</span> <span class="font-semibold text-brand-orange">បឹងទន្លេសាប</span> និងវាលស្រែសម្បូរបែប ប៉ះពាល់ដល់ជីវិតរៀងរាល់ថ្ងៃរបស់ប្រជាជនជាង ១៧ លាននាក់។',
            '<span class="font-semibold text-brand-orange">ភ្នំពេញ</span> ទីក្រុងរដ្ឋធានី ស្ថិតនៅចំណុចជួបរវាងទន្លេមេគង្គ និងទន្លេសាប។ ចាប់ពីទីនេះ <span class="font-semibold text-brand-orange">ព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ</span> ចែករំលែកសេចក្តីសង្ឃឹមក្នុងព្រះយេស៊ូវ ទៅកាន់ទីក្រុងដែលមានផ្សារ កាហ្វេ សាកលវិទ្យាល័យ និងសង្កាត់ដែលតំណាងគ្រប់ស្រទាប់សង្គម។',
            'មរតកនៃ <span class="font-semibold text-brand-orange">ចក្រខ្មែរ</span> ដោយពិសេស <span class="font-semibold text-brand-orange">អង្គរវត្ត</span> នៅតែលើកទឹកចិត្តជាតិ។ ប៉ុន្តែប្រវត្តិសាស្ត្រថ្មីៗមានទុក្ខលំបាកពីរបប <span class="font-semibold text-brand-orange">ខ្មែរក្រហម</span>។ សព្វថ្ងៃ កម្ពុជាកំពុងស្ដារឡើងវិញតាមការអប់រំ សហគ្រិនភាព និងជំនាន់ថ្មីដែលចង់បានសេចក្តីពិត ការស្ដារឡើងវិញ និងគោលបំណង។',
            '<span class="font-semibold text-brand-orange">សេវាកម្មនៅកម្ពុជា</span> កើតឡើងដោយការស្តាប់យ៉ាងល្អ គោរពវប្បធម៌ វិនិយោគក្នុងទំនាក់ទំនង និងបង្ហាញ <span class="font-semibold text-brand-orange">សារព្រះ</span> ដោយស្មោះត្រង់ជាភាសាខ្មែរ។ ចាប់ពីក្រុមផ្ទះដល់ការចេញផ្សាយក្នុងទីក្រុង យើងបំណងបំពាក់អ្នកជឿ និងបម្រើភ្នំពេញដោយភាពទាប និងសេចក្តីសង្ឃឹម។',
          ],
          cta: "ស្វែងយល់បន្ថែម",
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );

    return {
      hero: computed(() => content.value.hero),
      about: computed(() => content.value.about),
    };
  },
};
</script>
