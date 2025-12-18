<template>
  <div>


    <!-- Hero Section -->
    <section class="py-24 bg-gradient-to-br from-brand-blue via-brand-blue/80 to-brand-orange/90 text-white border-b border-white/10">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6">
        <p class="text-sm uppercase tracking-[0.4em] text-white/70">{{ hero.eyebrow }}</p>
        <h1 class="text-2xl md:text-3xl lg:text-4xl font-bold leading-tight">
          {{ hero.title }}
        </h1>
        <p class="text-lg md:text-xl text-white/85 font-light max-w-3xl mx-auto">
          {{ hero.subtitle }}
        </p>
      </div>
    </section>

    <!-- Leaders Section -->
    <section class="py-24 bg-gray-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid gap-10 md:grid-cols-2 xl:grid-cols-3">
          <article
            v-for="leader in leaders"
            :key="leader.name"
            class="group relative overflow-hidden rounded-[2.3rem] border border-white/15 bg-white/10 shadow-[0_20px_50px_rgba(15,23,42,0.18)] backdrop-blur-xl p-8 text-center flex flex-col items-center"
          >
            <div class="pointer-events-none absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-500">
              <div class="absolute inset-0 bg-gradient-to-br from-brand-blue/30 via-transparent to-brand-orange/40"></div>
            </div>
            <div class="relative flex flex-col items-center flex-1">
              <div class="relative mb-6">
                <div class="absolute inset-0 rounded-full bg-brand-orange/40 blur-3xl opacity-0 group-hover:opacity-60 transition-opacity duration-500"></div>
                <div class="mx-auto h-28 w-28 overflow-hidden rounded-full border-4 border-white/80 shadow-[0_10px_35px_rgba(15,23,42,0.3)]">
                  <img
                    :src="leader.photo"
                    :alt="`${leader.name} portrait`"
                    class="h-full w-full object-cover"
                  />
                </div>
              </div>
              <h3 class="text-xl font-semibold text-gray-900">{{ leader.name }}</h3>
              <p class="mt-2 text-brand-blue font-medium tracking-[0.3em] uppercase text-[0.65rem]">{{ leader.role }}</p>
              <p class="mt-3 text-sm text-gray-700 leading-relaxed max-h-32 overflow-hidden">
                {{ leader.bio }}
              </p>
            </div>
            <button
              type="button"
              class="relative mt-6 inline-flex items-center gap-2 rounded-full border border-brand-blue/40 px-5 py-2 text-[0.65rem] font-semibold tracking-[0.3em] uppercase text-brand-blue transition-all duration-300 hover:bg-brand-blue hover:text-white"
            >
              {{ labels.more }}
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </button>
          </article>
        </div>

        <div class="mt-16">
          <div class="text-center mb-10">
            <p class="text-sm uppercase tracking-[0.3em] text-gray-500">{{ team.eyebrow }}</p>
            <h3 class="text-2xl font-semibold text-gray-900">{{ team.title }}</h3>
          </div>
          <div class="grid gap-8 sm:grid-cols-2 lg:grid-cols-3">
            <article
              v-for="member in team.members"
              :key="member.name"
              class="rounded-3xl border border-gray-100 bg-white shadow-lg p-6 text-center flex flex-col items-center gap-3"
            >
              <div class="h-16 w-16 rounded-full bg-brand-blue/5 border border-brand-blue/15 flex items-center justify-center shadow-inner">
                <img :src="member.photo" :alt="`${member.name} icon`" class="h-8 w-8 object-contain" />
              </div>
              <div>
                <p class="text-base font-semibold text-gray-900">{{ member.name }}</p>
                <p class="text-sm uppercase tracking-[0.25em] text-brand-blue/80">{{ member.role }}</p>
              </div>
            </article>
          </div>
        </div>

        <div class="text-center mt-16">
          <p class="text-gray-600 mb-3 text-sm uppercase tracking-[0.3em]">{{ cta.eyebrow }}</p>
          <h4 class="text-2xl font-semibold text-gray-900 mb-6">{{ cta.title }}</h4>
          <router-link
            to="/contact"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors shadow-md"
          >
            {{ cta.button }}
            <svg class="ml-2 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </router-link>
        </div>
      </div>
    </section>

  </div>
</template>

<script>
import { computed } from 'vue'
import { storeToRefs } from 'pinia'
import { useLanguageStore } from '@/stores/language'
import rithPhoto from '@/assets/leadership/rith.jpg'
import longPhoto from '@/assets/leadership/long.jpeg'
import randyPhoto from '@/assets/leadership/randy.jpg'
import vuthaPhoto from '@/assets/leadership/vutha.jpg'
import curtisPhoto from '@/assets/leadership/curtis.jpg'
import teamIcon from '@/assets/leadership/team-icon.svg'

export default {
  name: 'Leadership',
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore())

    const translations = {
      en: {
        hero: {
          eyebrow: 'Leadership',
          title: 'Shepherding Phnom Penh with prayer, Scripture, and presence.',
          subtitle:
            'Our elders and ministry leads teach verse by verse, disciple across generations, and keep the church tethered to Jesus and His mission.'
        },
        labels: {
          more: 'Know More'
        },
        cta: {
          eyebrow: 'Next steps',
          title: 'Connect with our leadership team',
          button: 'Contact Us'
        },
        team: {
          eyebrow: 'Team Leaders',
          title: 'Serving across ministries every week',
          members: [
            { name: 'Kuyeng Khat', role: 'Childrens Ministry', photo: teamIcon },
            { name: 'Channy Saing', role: 'Youth Group', photo: teamIcon },
            { name: 'Daro Chim', role: 'Tech Team', photo: teamIcon },
            { name: 'Curtis Johnson', role: 'Worship', photo: teamIcon },
            { name: 'Sokhom Norn', role: 'Welcome Team', photo: teamIcon },
            { name: 'Savary Sang', role: 'Areyksat Outreach', photo: teamIcon },
            { name: 'Chanthy Chom', role: 'Accounting', photo: teamIcon },
            { name: 'Eva-Maria Sang', role: 'Administration', photo: teamIcon },
            { name: 'Vutha Ven', role: 'Translation', photo: teamIcon },
            { name: 'Long Sann', role: 'Facilities Service', photo: teamIcon }
          ]
        },
        leaders: [
          {
            name: 'Rith Sang',
            role: 'Lead Pastor',
            photo: rithPhoto,
            bio: 'Leads Calvary Chapel Phnom Penh with steady Bible teaching, prayer, and pastoral care that keeps the church centered on Jesus.'
          },
          {
            name: 'Long Sann',
            role: 'Pastor',
            photo: longPhoto,
            bio: 'Shepherds families and outreach teams, walking with people in discipleship and teaching Scripture in everyday life.'
          },
          {
            name: 'Vutha Ven',
            role: 'Pastor',
            photo: vuthaPhoto,
            bio: 'Guides worship and prayer gatherings, encouraging the church to seek God together and serve neighbors across the city.'
          },
          {
            name: 'Randy Fleming',
            role: 'Elder',
            photo: randyPhoto,
            bio: 'Serves as an elder with wise counsel, supporting teaching, pastoral care, and hospitality for the church family.'
          },
          {
            name: 'Curtis Johnson',
            role: 'Elder',
            photo: curtisPhoto,
            bio: 'Elder focused on encouragement and prayer, helping leaders and teams stay aligned to the gospel and mission.'
          }
        ]
      },
      kh: {
        hero: {
          eyebrow: 'ក្រុមភាពជាអ្នកដឹកនាំ',
          title: 'ថែរក្សាទីក្រុងភ្នំពេញដោយការអធិស្ឋាន ព្រះវចនៈ និងការគាំទ្រជិតស្និទ្ធ។',
          subtitle:
            'ជាន់ខ្ពស់ និងអ្នកដឹកនាំសេវាកម្ម បង្រៀនតាមជួរ បណ្តុះសិស្សគ្រប់ជំនាន់ និងរក្សាគ្រឹស្តសាសនាចក្រតភ្ជាប់នឹងព្រះយេស៊ូវ និងភារកិច្ចរបស់ទ្រង់។'
        },
        labels: {
          more: 'មើលបន្ថែម'
        },
        cta: {
          eyebrow: 'ជំហានបន្ទាប់',
          title: 'ទាក់ទងជាមួយក្រុមភាពជាអ្នកដឹកនាំ',
          button: 'ទាក់ទងមកកាន់ពួកយើង'
        },
        team: {
          eyebrow: 'Team Leaders',
          title: 'Serving across ministries every week',
          members: [
            { name: 'Kuyeng Khat', role: 'Childrens Ministry', photo: teamIcon },
            { name: 'Channy Saing', role: 'Youth Group', photo: teamIcon },
            { name: 'Daro Chim', role: 'Tech Team', photo: teamIcon },
            { name: 'Curtis Johnson', role: 'Worship', photo: teamIcon },
            { name: 'Sokhom Norn', role: 'Welcome Team', photo: teamIcon },
            { name: 'Savary Sang', role: 'Areyksat Outreach', photo: teamIcon },
            { name: 'Chanthy Chom', role: 'Accounting', photo: teamIcon },
            { name: 'Eva-Maria Sang', role: 'Administration', photo: teamIcon },
            { name: 'Vutha Ven', role: 'Translation', photo: teamIcon },
            { name: 'Long Sann', role: 'Facilities Service', photo: teamIcon }
          ]
        },
        leaders: [
          {
            name: 'Rith Sang',
            role: 'Lead Pastor',
            photo: rithPhoto,
            bio: 'ដឹកនាំ Calvary Chapel Phnom Penh ជាមួយការបង្រៀនព្រះគម្ពីរ ការអធិស្ឋាន និងការថែរក្សាសមាជិក ដើម្បីរក្សាសាសនាចក្រ​កណ្ដាលទៅលើព្រះយេស៊ូវ។'
          },
          {
            name: 'Long Sann',
            role: 'Pastor',
            photo: longPhoto,
            bio: 'ថែរក្សាគ្រួសារ និងក្រុមចេញផ្សាយ ដើរជាមួយមនុស្សក្នុងការបណ្តុះសិស្ស និងបង្រៀនព្រះវចនៈក្នុងជីវិតប្រចាំថ្ងៃ។'
          },
          {
            name: 'Vutha Ven',
            role: 'Pastor',
            photo: vuthaPhoto,
            bio: 'ណែនាំការថ្វាយបង្គំ និងការអធិស្ឋាន លើកទឹកចិត្តឲ្យសាសនាចក្រ​ស្វែងរកព្រះជាមួយគ្នា និងបម្រើអ្នកជិតខាងទូទាំងទីក្រុង។'
          },
          {
            name: 'Randy Fleming',
            role: 'Elder',
            photo: randyPhoto,
            bio: 'បម្រើជាអធិការក្រុមជាមួយប្រាជ្ញា គាំទ្រការបង្រៀន ការថែរក្សា និងការស្វាគមន៍សមាជិក។'
          },
          {
            name: 'Curtis Johnson',
            role: 'Elder',
            photo: curtisPhoto,
            bio: 'ផ្តោតលើការលើកទឹកចិត្ត និងការអធិស្ឋាន ជួយក្រុមដឹកនាំឲ្យស្ថិតស្របនឹងព្រះវចនៈ និងបេសកកម្ម។'
          }
        ]
      }
    }

    const content = computed(() => translations[selectedLanguage.value] || translations.en)

    return {
      hero: computed(() => content.value.hero),
      labels: computed(() => content.value.labels),
      cta: computed(() => content.value.cta),
      team: computed(() => content.value.team),
      leaders: computed(() => content.value.leaders)
    }
  }
}
</script>

