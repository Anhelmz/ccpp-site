<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[65vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-16"
    >
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="{
            backgroundImage: `linear-gradient(120deg, rgba(17,39,84,0.35), rgba(17,39,84,0.55)), url(${beliefHeroImage})`,
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
                circle at 18% 30%,
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
        class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6"
      >
        <p class="text-sm uppercase tracking-[0.35em] text-white/70">
          {{ hero.eyebrow }}
        </p>
        <h1
          class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight"
        >
          {{ hero.title }}
        </h1>
        <p class="text-lg md:text-xl text-white/85 font-light">
          {{ hero.subtitle }}
        </p>
        <div class="flex flex-wrap justify-center gap-3">
          <span
            v-for="point in heroPoints"
            :key="point"
            class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/80 text-xs uppercase tracking-[0.25em] backdrop-blur-sm"
          >
            {{ point }}
          </span>
        </div>
      </div>
    </section>

    <!-- Detailed Statements -->
    <section class="relative py-24 bg-brand-blue text-white overflow-hidden">
      <div class="absolute inset-0 opacity-80" aria-hidden="true">
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
      <div class="relative max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-12">
        <div class="text-center">
          <p class="text-sm uppercase tracking-[0.35em] text-white/70 mb-3">
            {{ statementsHeader.eyebrow }}
          </p>
          <h2 class="text-4xl md:text-5xl font-bold mb-4">
            {{ statementsHeader.title }}
          </h2>
          <p class="text-white/80 max-w-3xl mx-auto">
            {{ statementsHeader.description }}
          </p>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <article
            v-for="(statement, index) in detailedStatements"
            :key="statement.id"
            class="rounded-3xl border border-white/15 bg-white/5 backdrop-blur-lg p-8 space-y-4"
          >
            <div class="flex items-center gap-4">
              <span
                class="inline-flex h-12 w-12 items-center justify-center rounded-2xl bg-white/10 text-lg font-semibold"
              >
                {{ (index + 1).toString().padStart(2, "0") }}
              </span>
              <h3 class="text-xl font-semibold leading-snug">
                {{ statement.title }}
              </h3>
            </div>
            <p class="text-base leading-relaxed text-white/85">
              {{ statement.body }}
            </p>
            <p
              v-if="statement.reference"
              class="text-sm uppercase tracking-[0.3em] text-white/60"
            >
              {{ statement.reference }}
            </p>
          </article>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="py-20 bg-white">
      <div class="max-w-3xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
          {{ cta.eyebrow }}
        </p>
        <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
          {{ cta.title }}
        </h2>
        <p class="text-lg text-gray-600">
          {{ cta.description }}
        </p>
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors shadow-lg shadow-brand-orange/30"
          >
            {{ cta.primary }}
          </router-link>
          <router-link
            to="/about"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-blue/10 text-brand-blue font-semibold rounded-full border border-brand-blue/20 hover:bg-brand-blue/20 transition-colors"
          >
            {{ cta.secondary }}
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import beliefHeroImage from "@/assets/background/bg-hero3.jpg";

export default {
  name: "WhatWeBelieve",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          eyebrow: "Statement of faith",
          title: "What We Believe",
          subtitle:
            "Our doctrine is historic, Christ-centered, and fuels the way we teach, serve, and love Phnom Penh.",
        },
        heroPoints: [
          "Scripture",
          "Spirit-led",
          "Christ-centered",
          "Hope-filled",
        ],
        statementsHeader: {
          eyebrow: "Full statement",
          title: "We believe…",
          description:
            "The following articles summarize the doctrines we joyfully teach. Each one shapes the way we gather, disciple, and serve Phnom Penh.",
        },
        statements: [
          {
            id: "triune-god",
            title: "The Triune God",
            body: "That there is one living and true God, eternally existing in three persons, the Father, the Son, and the Holy Spirit, equal in power and glory; that this triune God created all, upholds all, and governs all.",
          },
          {
            id: "scripture",
            title: "Authority of Scripture",
            body: "That the Scriptures of the Old and New Testaments are the Word of God, fully inspired without error in the original manuscripts, and the infallible rule of faith and practice.",
          },
          {
            id: "father",
            title: "God the Father",
            body: "In God the Father, an infinite, personal Spirit, perfect in holiness, wisdom, power and love; that He concerns Himself mercifully in the affairs of men, hears and answers prayer, and saves from sin and death all who come to Him through Jesus Christ.",
          },
          {
            id: "jesus",
            title: "Jesus Christ the Son",
            body: "In Jesus Christ, God's only begotten Son, conceived by the Holy Spirit. We affirm His virgin birth, sinless life, miracles and teachings, His substitutionary atoning death, bodily resurrection, ascension into heaven, perpetual intercession for His people, and personal, visible return to earth.",
          },
          {
            id: "blessed-hope",
            title: "The Blessed Hope",
            body: "In the Blessed Hope, the personal, visible, pre-tribulation return of our Lord and Savior, Jesus Christ.",
          },
          {
            id: "resurrection",
            title: "Resurrection & Eternity",
            body: "In the bodily resurrection of the just and the unjust, the everlasting blessedness of the saved and the everlasting conscious punishment of the lost.",
          },
          {
            id: "holy-spirit",
            title: "The Holy Spirit",
            body: "In the Holy Spirit, who came forth from the Father and Son to convict the world of sin, righteousness, and judgment and to regenerate, sanctify and empower for ministry all who believe in Christ. We believe the Holy Spirit indwells every believer in Jesus Christ and that He is an abiding Helper, Teacher and Guide, active today in all His gifts.",
          },
          {
            id: "salvation",
            title: "Salvation & New Life",
            body: "That all men are sinners by nature and choice and therefore under condemnation, that God regenerates by the Holy Spirit those who repent of their sins and confess Jesus Christ as Lord; that Jesus Christ baptizes the seeking believer with the Holy Spirit and power for service, often subsequent to regeneration.",
          },
          {
            id: "church",
            title: "The Church",
            body: "In the universal Church, the living spiritual body of which Christ is the Head and all regenerated persons are members.",
          },
          {
            id: "ordinances",
            title: "Baptism & Communion",
            body: "That the Lord Jesus Christ committed two ordinances to the Church: baptism and the Lord's Supper. We practice baptism by immersion and celebrate communion open to all believers.",
          },
          {
            id: "kingdom",
            title: "Christ’s Kingdom",
            body: "In the personal, visible return of Christ to earth and the establishment of His Kingdom; in the resurrection of the body, the final judgment and eternal blessing of the righteous and endless separation of the wicked.",
          },
          {
            id: "discernment",
            title: "Discernment by the Word",
            body: "That the Bible is the inspired Word of God, complete and entire, through which the Holy Spirit gives authority to test all things and discern what is of God.",
            reference: "1 Corinthians 2:13",
          },
        ],
        cta: {
          eyebrow: "Connect",
          title: "Want to talk through these beliefs?",
          description:
            "Our pastors and team would love to open Scripture with you, answer questions, and help you take the next step with Jesus.",
          primary: "Contact Us",
          secondary: "Discover Calvary Chapel",
        },
      },
      kh: {
        hero: {
          eyebrow: "សេចក្តីប្រកាសជំនឿ",
          title: "អ្វីដែលយើងជឿ",
          subtitle:
            "សេចក្តីបង្រៀនរបស់យើងមានប្រវត្តិសាស្ត្រ ដាក់ព្រះគ្រីស្ទជាកណ្ដាល និងជំរុញរបៀបបង្រៀន បម្រើ និងស្រឡាញ់ទីក្រុងភ្នំពេញ។",
        },
        heroPoints: [
          "ព្រះវចនៈ",
          "ដឹកនាំដោយព្រះវិញ្ញាណ",
          "កណ្ដាលគ្រីស្ទ",
          "ពោរពេញដោយសេចក្តីសង្ឃឹម",
        ],
        statementsHeader: {
          eyebrow: "សេចក្តីប្រកាសពេញលេញ",
          title: "យើងជឿថា…",
          description:
            "អត្ថបទខាងក្រោមសង្ខេបសេចក្តីបង្រៀនដែលយើងរៀបរាប់ដោយអស់ពីចិត្ត។ អ្វីៗទាំងនេះប៉ះពាល់ដល់របៀបដែលយើងជួបជុំ បណ្តុះសិស្ស និងបម្រើទីក្រុងភ្នំពេញ។",
        },
        statements: [
          {
            id: "triune-god",
            title: "ព្រះត្រីឯក",
            body: "មានព្រះមួយអង្គដ៏រស់ និងពិត ស្ថិតជារៀងរហូតក្នុងបីបុគ្គល គឺព្រះវរបិតា ព្រះវរបុត្រា និងព្រះវិញ្ញាណបរិសុទ្ធ ស្មើគ្នាអំណាច និងសិរីល្អ; ព្រះត្រីឯកទ្រង់បានបង្កើត សង្រ្គោះ និងគ្រប់គ្រងអស់ទាំងអស់។",
          },
          {
            id: "scripture",
            title: "អធិការកិច្ចនៃព្រះគម្ពីរ",
            body: "ព្រះគម្ពីរសញ្ញាចាស់ និងសញ្ញាថ្មីជាព្រះបន្ទូលរបស់ព្រះ ទទួលបានបន្ទាល់ពេញលេញ គ្មានកំហុសក្នុងអត្ថន័យដើម ហើយជាក្បួនជំនឿ និងអនុវត្តដែលមិនខ្វះខាត។",
          },
          {
            id: "father",
            title: "ព្រះវរបិតា",
            body: "យើងជឿលើព្រះវរបិតា ព្រះវិញ្ញាណផ្ទាល់ខ្លួនអស់កល្ប បរិសុទ្ធ ប្រាជ្ញា អំណាច និងសេចក្តីស្រឡាញ់; ទ្រង់មានមេត្តាដល់រឿងរបារបស់មនុស្ស ស្តាប់ និងឆ្លើយអធិស្ឋាន និងសង្គ្រោះអ្នកដែលមករកទ្រង់តាមព្រះយេស៊ូវគ្រីស្ទពីបាប និងសេចក្ដីស្លាប់។",
          },
          {
            id: "jesus",
            title: "ព្រះយេស៊ូវគ្រីស្ទ ព្រះរាជបុត្រា",
            body: "យើងជឿថា ព្រះយេស៊ូវគ្រីស្ទ ជាព្រះរាជបុត្រាតែមួយ របស់ព្រះ គង្វាលដោយព្រះវិញ្ញាណបរិសុទ្ធ កំណើតពីក្រមុំ កំពុងរស់គ្មានបាប អនុញ្ញាតអស្ចារ្យ និងការបង្រៀន រងទុក្ខជំនួសយើង លើឈើឆ្កាង រាងកាយរស់ឡើងវិញ ឡើងទៅស្ថានសួគ៌ អធិស្ឋានជានិច្ចសម្រាប់ប្រជាជនរបស់ទ្រង់ និងនឹងត្រឡប់មកវិញផ្ទាល់មុខ។",
          },
          {
            id: "blessed-hope",
            title: "សេចក្តីសង្ឃឹមអស្ចារ្យ",
            body: "យើងជឿសេចក្តីសង្ឃឹមអស្ចារ្យ គឺការត្រឡប់មកវិញផ្ទាល់មុខ មើលឃើញ មុនពេលទុក្ខលំបាករបស់ព្រះអម្ចាស់ និងព្រះជាម្ចាស់យើង ព្រះយេស៊ូវគ្រីស្ទ។",
          },
          {
            id: "resurrection",
            title: "ការរស់ឡើងវិញ និងអស់កល្ប",
            body: "យើងជឿការរស់ឡើងវិញជារូបកាយរបស់អ្នក義 និងអ្នកដែលមិន义 ការអំណរអស់កល្បរបស់អ្នកបានសង្គ្រោះ និងការផាកពិន័យដ៏មានការយល់ដឹងអស់កល្បរបស់អ្នកបាត់បង់។",
          },
          {
            id: "holy-spirit",
            title: "ព្រះវិញ្ញាណបរិសុទ្ធ",
            body: "យើងជឿថា ព្រះវិញ្ញាណបរិសុទ្ធ បានមកពីព្រះវរបិតា និងព្រះវរបុត្រា ដើម្បីទាស់បាប ប្រាថ្នាភាព義 និងការកាត់ក្តី ហើយបង្កើតថ្មី បរិសុទ្ធ និងប្រទានអំណាចសម្រាប់សេវាកម្មដល់អ្នកជឿទាំងអស់។ ព្រះវិញ្ញាណស្នាក់នៅក្នុងអ្នកជឿទាំងអស់ ជាជំនួយការ គ្រូ និងអ្នកណែនាំជានិច្ច សកម្មក្នុងអំណោយទាំងអស់នៅថ្ងៃនេះ។",
          },
          {
            id: "salvation",
            title: "សេចក្តីសង្គ្រោះ និងជីវិតថ្មី",
            body: "មនុស្សទាំងអស់ជាអ្នកបាបតាមធម្មជាតិ និងការជ្រើសរើស ហើយស្ថិតក្រោមទោស; ព្រះអង្គបង្កើតថ្មីតាមព្រះវិញ្ញាណបរិសុទ្ធដល់អ្នកដែលប្រែចិត្តពីបាប និងសារភាពព្រះយេស៊ូវគ្រីស្ទជាព្រះអម្ចាស់; ព្រះយេស៊ូវជ្រមុជអ្នកជឿដែលស្វែងរក ដោយព្រះវិញ្ញាណបរិសុទ្ធ និងអំណាចសម្រាប់សេវាកម្ម ជាញឹកញាប់បន្ទាប់ពីការចេញកំណើតថ្មី។",
          },
          {
            id: "church",
            title: "សាសនាចក្រគ្រីស្ទ",
            body: "យើងជឿសាសនាចក្រសកល ជារូបកាយវិញ្ញាណរស់ ដែលគ្រីស្ទជាប្រមុខ និងអ្នកកើតថ្មីទាំងអស់ជាសមាជិក។",
          },
          {
            id: "ordinances",
            title: "បុណ្យជ្រមុជទឹក និងបញ្ញាក់ព្រះអម្ចាស់",
            body: "ព្រះយេស៊ូវប្រគល់ពិធីបុណ្យពីរដល់សាសនាចក្រ គឺបុណ្យជ្រមុជទឹក និងអាហារព្រះអម្ចាស់។ យើងអនុវត្តបុណ្យជ្រមុជទឹកដោយដាក់ជ្រមុជ ហើយអញ្ជើញអ្នកជឿទាំងអស់ចូលរួមអាហារព្រះអម្ចាស់។",
          },
          {
            id: "kingdom",
            title: "ព្រះរាជាណាចក្ររបស់គ្រីស្ទ",
            body: "យើងជឿការត្រឡប់មកផ្ទាល់មុខរបស់គ្រីស្ទមកផែនដី និងការបង្កើតព្រះរាជាណាចក្ររបស់ទ្រង់; ការរស់ឡើងវិញនៃរាងកាយ ការកាត់ក្តីចុងក្រោយ ការអំណរអស់កល្បរបស់អ្នក義 និងការបែកចែកអស់កល្បរបស់អ្នកមិន義។",
          },
          {
            id: "discernment",
            title: "ការរក្សាចិត្តដោយព្រះបន្ទូល",
            body: "ព្រះគម្ពីរ ជាព្រះបន្ទូលដែលទទួលដកដង្ហើមពីព្រះ គ្រប់គ្រាន់ ដែលព្រះវិញ្ញាណបរិសុទ្ធប្រទានអធិការកិច្ចដើម្បីសាកល្បងអ្វីៗ និងសម្គាល់អ្វីដែលមកពីព្រះ។",
            reference: "1 កូរិនថូស 2:13",
          },
        ],
        cta: {
          eyebrow: "តភ្ជាប់",
          title: "ចង់ពិភាក្សាអំពីជំនឿទាំងនេះទេ?",
          description:
            "គ្រូគង្វាល និងក្រុមចង់បើកព្រះវចនៈជាមួយអ្នក ឆ្លើយសំណួរ និងជួយអ្នកធ្វើជំហានបន្ទាប់ជាមួយព្រះយេស៊ូវ។",
          primary: "ទាក់ទងមកកាន់ពួកយើង",
          secondary: "ស្វែងយល់អំពី Calvary Chapel",
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );

    return {
      beliefHeroImage,
      hero: computed(() => content.value.hero),
      heroPoints: computed(() => content.value.heroPoints),
      statementsHeader: computed(() => content.value.statementsHeader),
      detailedStatements: computed(() => content.value.statements),
      cta: computed(() => content.value.cta),
    };
  },
};
</script>
