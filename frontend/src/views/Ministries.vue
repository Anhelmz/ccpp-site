<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[45vh] flex items-center overflow-hidden pt-16 md:pt-20 pb-12 bg-gradient-to-br from-[#111d33] via-[#1d2f48] to-[#213d5c] text-white"
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
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <section class="text-center space-y-6">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
            {{ intro.eyebrow }}
          </p>
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
            {{ intro.title }}
          </h2>
          <div
            class="space-y-4 text-left text-base md:text-lg text-gray-900 leading-relaxed max-w-3xl mx-auto"
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
            v-for="card in cards"
            :key="card.id"
            class="group relative rounded-[32px] border border-gray-100 bg-white shadow-2xl overflow-hidden flex flex-col"
          >
            <div class="relative h-52 overflow-hidden">
              <img
                :src="card.image"
                :alt="card.title"
                class="w-full h-full object-cover transition-transform duration-700 group-hover:scale-105"
              />
              <div
                class="absolute inset-0 bg-gradient-to-t from-black/60 via-transparent to-transparent"
              ></div>
              <span
                class="absolute top-4 left-4 inline-flex items-center px-3 py-1 rounded-full text-xs uppercase tracking-[0.35em] bg-white/90 text-gray-900"
              >
                {{ card.tag }}
              </span>
            </div>
            <div class="p-8 space-y-4 flex flex-col flex-1">
              <h3 class="text-2xl font-semibold text-gray-900">
                {{ card.title }}
              </h3>
              <p class="text-gray-600 flex-1">
                {{ card.body }}
              </p>
              <router-link
                :to="card.to"
                class="inline-flex items-center justify-center px-4 py-2 text-base font-semibold text-white bg-brand-orange rounded-full shadow-lg shadow-brand-orange/30 transition-all duration-300 group-hover:-translate-y-0.5 group-hover:bg-brand-orange/90"
              >
                {{ card.cta }}
                <svg
                  class="ml-2 w-4 h-4"
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
              </router-link>
            </div>
          </article>
        </div>

        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-lg hover:bg-brand-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            {{ cta.primary }}
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import youthMinistryImage from "@/assets/gallery/571314951_1275851217909869_307313164703676157_n.jpg";
import graceChurchImage from "@/assets/gallery/548188348_1236675751827416_5183731636046995004_n.jpg";
import outreachesImage from "@/assets/gallery/514246564_1275851177909873_6093560206728412070_n.jpg";

export default {
  name: "Ministries",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          title: "Ministries",
          subtitle:
            "Serving Phnom Penh with Scripture, worship, discipleship, and practical love.",
          chips: ["Discipleship", "Worship", "Families", "Outreaches"],
        },
        intro: {
          eyebrow: "Question",
          title: "What are our ministries?",
          body: [
            'Each ministry exists to keep <span class="font-semibold text-brand-orange">Jesus at the center</span>, meeting people in Phnom Penh with Scripture, Spirit-led care, and authentic community.',
            'We disciple kids, students, adults, and families through <span class="font-semibold text-brand-orange">Bible study</span>, <span class="font-semibold text-brand-orange">worship</span>, and <span class="font-semibold text-brand-orange">prayer</span>, while practical outreaches serve neighborhoods with compassion.',
            "Whether you’re new to faith or ready to lead, there is a place for you to grow, belong, and serve alongside Calvary Chapel Phnom Penh.",
          ],
        },
        cards: [
          {
            id: "youth",
            tag: "Youth",
            title: "Youth Ministry",
            body: "Teens gather weekly for worship, teaching, and discipleship that equips them to follow Jesus boldly and love Phnom Penh well.",
            cta: "Learn More",
            to: "/ministries/youth-ministry",
            image: youthMinistryImage,
          },
          {
            id: "grace",
            tag: "Grace",
            title: "Grace Church",
            body: "A network of house gatherings and Sunday services proclaiming the grace of God and forming communities that practice forgiveness and hospitality.",
            cta: "Learn More",
            to: "/ministries/grace-church",
            image: graceChurchImage,
          },
          {
            id: "outreach",
            tag: "Outreach",
            title: "Outreaches",
            body: "Teams bring meals, prayer, medical care, and encouragement into neighborhoods, pointing friends and families to the hope of Jesus.",
            cta: "Learn More",
            to: "/ministries/outreaches",
            image: outreachesImage,
          },
        ],
        cta: {
          primary: "Contact Us",
        },
      },
      kh: {
        hero: {
          title: "សេវាកម្ម",
          subtitle:
            "បម្រើទីក្រុងភ្នំពេញដោយព្រះវចនៈ ការថ្វាយបង្គំ ការបណ្តុះសិស្ស និងសេចក្តីស្រឡាញ់ជាក់ស្តែង។",
          chips: ["ការបណ្តុះសិស្ស", "ការថ្វាយបង្គំ", "គ្រួសារ", "ការចេញផ្សាយ"],
        },
        intro: {
          eyebrow: "សំណួរ",
          title: "សេវាកម្មរបស់យើងមានអ្វីខ្លះ?",
          body: [
            'សេវាកម្មនីមួយៗមានគោលបំណងរក្សា <span class="font-semibold text-brand-orange">ព្រះយេស៊ូវជាមជ្ឈមណ្ឌល</span> ជួបជាមួយមនុស្សនៅទីក្រុងភ្នំពេញតាមរយៈព្រះវចនៈ ការថែរក្សាដឹកនាំដោយព្រះវិញ្ញាណ និងសហគមន៍ពិតប្រាកដ។',
            'យើងបណ្តុះសិស្សកុមារ និស្សិត មនុស្សពេញវ័យ និងគ្រួសារ តាមរយៈ <span class="font-semibold text-brand-orange">ការសិក្សាព្រះគម្ពីរ</span> <span class="font-semibold text-brand-orange">ការថ្វាយបង្គំ</span> និង <span class="font-semibold text-brand-orange">ការអធិស្ឋាន</span> ខណៈដែលការចេញផ្សាយជាក់ស្តែងបម្រើសង្កាត់ដោយមេត្តាករុណា។',
            "មិនថាអ្នកថ្មីក្នុងជំនឿ ឬរួចរាល់ដើម្បីដឹកនាំ ក៏មានកន្លែងសម្រាប់អ្នកក្នុងការលូតលាស់ រួមចំណែក និងបម្រើជាមួយព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ។",
          ],
        },
        cards: [
          {
            id: "youth",
            tag: "យុវជន",
            title: "សេវាយុវជន",
            body: "យុវជនជួបជុំរៀងរាល់សប្ដាហ៍ សម្រាប់ការថ្វាយបង្គំ ការបង្រៀន និងការបណ្តុះសិស្ស ដែលបំពាក់ពួកគេឱ្យតាមព្រះយេស៊ូវដោយក្លាហាន និងស្រឡាញ់ទីក្រុងភ្នំពេញ។",
            cta: "ស្វែងយល់បន្ថែម",
            to: "/ministries/youth-ministry",
            image: youthMinistryImage,
          },
          {
            id: "grace",
            tag: "ព្រះគុណ",
            title: "សេវាព្រះគុណ",
            body: "បណ្តាញក្រុមផ្ទះ និងសេចក្តីប្រកាសថ្ងៃអាទិត្យ បង្ហាញព្រះគុណរបស់ព្រះ និងបង្កើតសហគមន៍ដែលអនុវត្តការអភ័យទោស និងការស្វាគមន៍។",
            cta: "ស្វែងយល់បន្ថែម",
            to: "/ministries/grace-church",
            image: graceChurchImage,
          },
          {
            id: "outreach",
            tag: "ចេញផ្សាយ",
            title: "ការចេញផ្សាយ",
            body: "ក្រុមនាំអាហារ ការអធិស្ឋាន ការថែទាំសុខភាព និងការលើកទឹកចិត្តចូលសង្កាត់ ដឹកនាំមិត្តភក្តិ និងគ្រួសារទៅរកសេចក្តីសង្ឃឹមក្នុងព្រះយេស៊ូវ។",
            cta: "ស្វែងយល់បន្ថែម",
            to: "/ministries/outreaches",
            image: outreachesImage,
          },
        ],
        cta: {
          primary: "ទាក់ទងមកកាន់ពួកយើង",
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );

    return {
      hero: computed(() => content.value.hero),
      intro: computed(() => content.value.intro),
      cards: computed(() => content.value.cards),
      cta: computed(() => content.value.cta),
    };
  },
};
</script>
