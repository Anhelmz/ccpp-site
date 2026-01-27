<template>
  <div>
    <!-- Hero Section -->
    <section
      class="relative min-h-[55vh] flex items-center overflow-hidden pt-20 md:pt-24 pb-16 bg-gradient-to-br from-[#0f1f33] via-[#182f4a] to-[#1f4061] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-25 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 20% 30%,
                rgba(255, 255, 255, 0.25),
                transparent 45%
              ),
              radial-gradient(
                circle at 80% 10%,
                rgba(255, 255, 255, 0.18),
                transparent 35%
              ),
              radial-gradient(
                circle at 60% 80%,
                rgba(255, 255, 255, 0.2),
                transparent 30%
              );
          "
        ></div>
        <div
          class="absolute inset-y-[-30%] left-[-10%] w-1/2 bg-gradient-to-r from-brand-orange/25 via-transparent to-transparent blur-[150px] opacity-60"
        ></div>
        <div
          class="absolute inset-y-[-40%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/35 via-transparent to-transparent blur-[200px] opacity-65"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/60 via-black/30 to-transparent"
        ></div>
      </div>
      <div class="relative z-10 w-full max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="grid md:grid-cols-2 gap-10 items-center">
          <div class="space-y-6 text-white">
            <div
              class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-white/10 border border-white/15 text-xs uppercase tracking-[0.25em]"
            >
              {{ hero.eyebrow }}
            </div>
            <div class="space-y-3">
              <h1
                class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight"
              >
                {{ hero.title }}
              </h1>
              <p class="text-lg md:text-xl text-white/85 max-w-2xl font-light">
                {{ hero.subtitle }}
              </p>
            </div>
            <div class="flex flex-wrap justify-start gap-3">
              <span
                v-for="chip in hero.chips"
                :key="chip"
                class="inline-flex items-center px-4 py-2 rounded-full border border-white/25 text-white/85 text-xs uppercase tracking-[0.3em] backdrop-blur-sm"
              >
                {{ chip }}
              </span>
            </div>
            <ul class="space-y-3 text-white/90">
              <li
                v-for="(item, idx) in hero.checklist"
                :key="idx"
                class="flex items-start gap-3"
              >
                <span
                  class="flex h-9 w-9 items-center justify-center rounded-full bg-white/10 border border-white/10"
                >
                  <svg
                    class="w-5 h-5"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M5 13l4 4L19 7"
                    />
                  </svg>
                </span>
                <div class="leading-relaxed">{{ item }}</div>
              </li>
            </ul>
          </div>
          <div
            class="bg-white/10 rounded-2xl p-6 md:p-7 border border-white/15 shadow-2xl shadow-black/20 backdrop-blur-md"
          >
            <p class="text-sm uppercase tracking-[0.25em] text-white/70 mb-4">
              {{ highlights.title }}
            </p>
            <div class="space-y-4">
              <div
                v-for="item in highlights.items"
                :key="item.title"
                class="flex items-start gap-4"
              >
                <span
                  class="flex h-10 w-10 items-center justify-center rounded-xl bg-white/10 border border-white/15"
                >
                  <svg
                    class="w-6 h-6 text-white"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                    />
                  </svg>
                </span>
                <div>
                  <p class="text-white font-semibold">{{ item.title }}</p>
                  <p class="text-white/80 text-sm">{{ item.detail }}</p>
                </div>
              </div>
            </div>
            <div
              class="mt-6 flex items-center justify-between text-white/80 text-sm border-t border-white/10 pt-4"
            >
              <span>{{ highlights.footer }}</span>
              <span class="inline-flex items-center gap-2 text-white">
                <span
                  class="h-2 w-2 rounded-full bg-brand-orange animate-pulse"
                ></span>
                {{ highlights.status }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Content Section -->
    <section class="bg-primary py-12 md:py-16">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 space-y-10">
        <section class="space-y-6">
          <div class="text-center space-y-3">
            <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
              {{ how.eyebrow }}
            </p>
            <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
              {{ how.title }}
            </h2>
            <p class="text-lg text-gray-700 max-w-4xl mx-auto leading-relaxed">
              {{ how.helper }}
            </p>
          </div>
          <div class="grid md:grid-cols-3 gap-4">
            <div
              v-for="(step, idx) in how.steps"
              :key="step.title"
              class="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 md:p-6 space-y-2"
            >
              <div
                class="inline-flex items-center justify-center h-9 w-9 rounded-full bg-main/10 text-main font-semibold"
              >
                {{ idx + 1 }}
              </div>
              <h3 class="text-lg font-semibold text-gray-900">
                {{ step.title }}
              </h3>
              <p class="text-gray-600 text-sm leading-relaxed">
                {{ step.body }}
              </p>
            </div>
          </div>
        </section>
        <!-- Loading State -->
        <div v-if="loading" class="py-12">
          <div class="text-center">
            <div
              class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-main mb-4"
            ></div>
            <p class="text-gray-600">{{ loadingText }}</p>
          </div>
        </div>

        <!-- Error State -->
        <div v-else-if="error" class="py-12 text-center">
          <p class="text-red-600">{{ error }}</p>
        </div>

        <!-- Calendar View -->
        <div v-else>
          <div
            class="bg-white/95 backdrop-blur-sm rounded-[32px] shadow-2xl p-6 md:p-8 border border-gray-100"
          >
            <!-- Calendar Header -->
            <div class="text-center mb-8 space-y-2">
              <h2 class="text-3xl md:text-4xl font-bold text-gray-900">
                {{ currentMonthName }} {{ currentYear }}
              </h2>
              <p class="text-gray-600 text-sm md:text-base">
                {{ calendarCopy.helper }}
              </p>
              <div
                class="flex justify-between items-center max-w-md mx-auto mt-4"
              >
                <button
                  class="p-2 text-main hover:bg-gray-100 rounded-lg transition-colors"
                  :aria-label="calendarCopy.prevAria"
                  @click="previousMonth"
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
                    ></path>
                  </svg>
                </button>
                <button
                  class="px-6 py-2 bg-main text-black font-semibold rounded-full hover:bg-main/90 transition-colors shadow-md shadow-main/30"
                  @click="resetToCurrentMonth"
                >
                  {{ calendarCopy.today }}
                </button>
                <button
                  class="p-2 text-main hover:bg-gray-100 rounded-lg transition-colors"
                  :aria-label="calendarCopy.nextAria"
                  @click="nextMonth"
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
                    ></path>
                  </svg>
                </button>
              </div>
            </div>

            <!-- Calendar Grid -->
            <div class="grid grid-cols-7 gap-2 mb-4">
              <div
                v-for="day in weekdayLabels"
                :key="day"
                class="text-center font-bold text-gray-700 py-3 text-sm md:text-base"
              >
                {{ day }}
              </div>
            </div>

            <div class="grid grid-cols-7 gap-2">
              <div
                v-for="day in calendarDays"
                :key="day.date"
                :class="[
                  'min-h-24 md:min-h-32 p-3 border rounded-2xl cursor-pointer transition-all hover:-translate-y-[1px] hover:shadow-md relative',
                  day.isToday
                    ? 'bg-main text-gray-900 border-main shadow-lg shadow-main/20'
                    : day.isCurrentMonth
                      ? 'bg-white hover:bg-gray-50 border-gray-200 hover:border-main/30'
                      : 'bg-gray-50 text-gray-400 border-gray-100 opacity-70',
                  day.hasEvents && !day.isToday
                    ? 'border-brand-orange/50 shadow-sm'
                    : '',
                ]"
                @click="selectDate(day)"
              >
                <div class="flex items-start justify-between mb-2">
                  <div
                    :class="[
                      'text-sm md:text-base font-semibold',
                      day.isToday
                        ? 'text-gray-900'
                        : day.isCurrentMonth
                          ? 'text-gray-900'
                          : 'text-gray-400',
                    ]"
                  >
                    {{ day.day }}
                  </div>
                  <div
                    v-if="day.events && day.events.length > 0"
                    class="text-[11px] font-semibold rounded-full px-2 py-0.5"
                    :class="
                      day.isToday
                        ? 'bg-gray-900/10 text-gray-900'
                        : 'bg-brand-orange/10 text-brand-orange'
                    "
                  >
                    {{ day.events.length }} {{ calendarCopy.eventCountLabel }}
                  </div>
                </div>
                <div
                  v-if="day.events && day.events.length > 0"
                  class="space-y-1"
                >
                  <div
                    v-for="event in day.events.slice(0, 2)"
                    :key="event.occurrenceId || event.id"
                    class="text-xs bg-brand-orange/15 text-brand-orange rounded-full px-2 py-1 truncate font-medium cursor-pointer hover:bg-brand-orange/25 transition-colors"
                    @click.stop="openModal(event)"
                  >
                    {{ event.title }}
                  </div>
                  <div
                    v-if="day.events.length > 2"
                    class="text-xs text-gray-500 font-medium cursor-pointer hover:text-gray-700 transition-colors"
                    @click.stop="selectDate(day)"
                  >
                    +{{ day.events.length - 2 }} {{ calendarCopy.moreLabel }}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Selected Date Events -->
          <div
            v-if="selectedDateEvents.length > 0"
            class="bg-white/95 rounded-[32px] shadow-2xl p-6 md:p-8 mt-8 border border-gray-100"
          >
            <h3 class="text-2xl md:text-3xl font-bold text-gray-900 mb-6">
              {{ calendarCopy.selectedPrefix }}
              {{ formatFullDate(selectedDay) }}
            </h3>
            <div class="space-y-3">
              <div
                v-for="event in selectedDateEvents"
                :key="event.occurrenceId || event.id"
                class="border-l-4 border-main pl-4 md:pl-6 py-3 md:py-4 hover:bg-gray-50 transition-all rounded-r-lg cursor-pointer group bg-white"
                @click="openModal(event)"
              >
                <h4
                  class="text-lg md:text-xl font-bold text-gray-900 mb-1 group-hover:text-main transition-colors"
                >
                  {{ event.title }}
                </h4>
                <div class="flex flex-wrap items-center gap-3 text-sm text-gray-600">
                  <span v-if="event.timeRange" class="font-medium">
                    {{ event.timeRange }}
                  </span>
                  <span v-if="event.location" class="font-medium">
                    {{ event.location }}
                  </span>
                  <span v-if="event.summary" class="text-gray-500 text-xs">
                    {{ event.summary }}
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Event Modal -->
    <div
      v-if="selectedEvent"
      class="fixed inset-0 bg-black bg-opacity-60 z-50 flex items-center justify-center p-4 backdrop-blur-sm"
      @click="closeModal"
    >
      <div
        class="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-hidden shadow-2xl transform transition-all"
        @click.stop
      >
        <!-- Modal Header -->
        <div
          class="sticky top-0 bg-white px-8 py-6 flex justify-between items-start z-10 border-b border-gray-200"
        >
          <div class="flex-1">
            <h2 class="text-3xl font-bold mb-4 pr-8" style="color: #000000 !important;">
              {{ selectedEvent.title }}
            </h2>
            <div class="flex flex-wrap items-center gap-4 text-sm">
              <div class="flex items-center px-4 py-2 rounded-full bg-blue-500 text-white">
                <svg
                  class="w-5 h-5 mr-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                  />
                </svg>
                <span class="font-medium">{{
                  formatFullDate(selectedEvent.startTime)
                }}</span>
              </div>
              <div
                v-if="selectedEvent.timeRange"
                class="flex items-center px-4 py-2 rounded-full bg-brand-orange text-white"
              >
                <svg
                  class="w-5 h-5 mr-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
                <span class="font-medium">{{ selectedEvent.timeRange }}</span>
              </div>
              <div
                v-if="selectedEvent.location"
                class="flex items-center px-4 py-2 rounded-full bg-green-500 text-white"
              >
                <svg
                  class="w-5 h-5 mr-2"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 11a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                </svg>
                <span class="font-medium">{{ selectedEvent.location }}</span>
              </div>
            </div>
          </div>
          <button
            class="p-2 hover:bg-gray-100 rounded-lg transition-colors"
            @click="closeModal"
            style="color: #000000 !important;"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              style="color: #000000 !important;"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M6 18L18 6M6 6l12 12"
              ></path>
            </svg>
          </button>
        </div>

        <!-- Modal Content -->
        <div class="px-8 py-8 overflow-y-auto max-h-[calc(90vh-180px)] bg-white" style="color: #000000 !important;">
          <div class="max-w-none" style="color: #000000 !important;">
            <div v-if="selectedEvent.summary" class="mb-6">
              <p class="text-lg font-medium leading-relaxed" style="color: #000000 !important;">
                {{ selectedEvent.summary }}
              </p>
            </div>
            <div v-if="selectedEvent.content">
              <p
                class="whitespace-pre-line leading-relaxed text-lg"
                style="color: #000000 !important;"
              >
                {{ selectedEvent.content }}
              </p>
            </div>
            <div v-else-if="selectedEvent.description">
              <p
                class="whitespace-pre-line leading-relaxed text-lg"
                style="color: #000000 !important;"
              >
                {{ selectedEvent.description }}
              </p>
            </div>
            <div v-else class="italic" style="color: #000000 !important;">
              No additional details available.
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed, watch } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import { eventService } from "@/services/eventService";

export default {
  name: "Calendar",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const loading = ref(true);
    const error = ref("");
    const events = ref([]);
    const selectedEvent = ref(null);
    // Use UTC for month/year to ensure consistency
    const currentMonth = ref(new Date().getUTCMonth());
    const currentYear = ref(new Date().getUTCFullYear());
    const selectedDay = ref(null);

    const translations = {
      en: {
        hero: {
          eyebrow: "Service Calendar",
          title: "Church Calendar",
          subtitle:
            "One place to see every worship gathering, prayer night, outreach, and family moment we have planned.",
          chips: ["Monthly Rhythm", "Worship & Prayer", "Youth & Families"],
          checklist: [
            "Browse by month and tap highlighted days.",
            "Get times, locations, and full details.",
            "Quickly find services, classes, and prayer nights.",
          ],
        },
        highlights: {
          title: "This month highlights",
          footer: "Live schedule · Up to date",
          status: "Live",
          items: [
            { title: "Sunday Worship", detail: "Every Sunday · 9:00 AM" },
            { title: "Youth Group", detail: "Every Saturday · 6:00 PM" },
            { title: "Prayer Night", detail: "First Saturday · 7:00 PM" },
          ],
        },
        how: {
          eyebrow: "How to use",
          title: "Use the calendar quickly",
          helper: "Select a highlighted day to see events and details.",
          steps: [
            {
              title: "Browse by month",
              body: "Use previous/next or Today to jump back to the current month.",
            },
            {
              title: "Pick highlighted days",
              body: "Highlighted dates show how many events are scheduled.",
            },
            {
              title: "Open an event",
              body: "See time, location, and description, then share with friends.",
            },
          ],
        },
        calendar: {
          helper: "Select a highlighted day to view events and details.",
          weekdays: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
          eventCountLabel: "events",
          moreLabel: "more",
          selectedPrefix: "Events on",
          today: "Today",
          prevAria: "Previous month",
          nextAria: "Next month",
        },
        loading: "Loading calendar...",
        error: "Failed to load calendar",
        months: [
          "January",
          "February",
          "March",
          "April",
          "May",
          "June",
          "July",
          "August",
          "September",
          "October",
          "November",
          "December",
        ],
        events: {
          sunday: {
            title: "Sunday Worship Service",
            summary:
              "Join us every Sunday for worship, prayer, and fellowship.",
            content: `We gather every Sunday morning for worship and teaching from God's Word.\n\nService Schedule:\n- Praise and Worship: 9:00 AM\n- Prayer: 9:30 AM\n- Message: 10:00 AM\n- Fellowship: 11:30 AM\n\nAll are welcome! We look forward to worshipping with you.`,
            time: "9:00 AM",
            location: "Main Sanctuary",
          },
          sundaySchool: {
            title: "Sunday School",
            summary: "Bible study classes for all ages.",
            content: `Join us for Sunday School classes for children, youth, and adults.\n\nAll classes run from 10:30 AM to 11:30 AM\nLocation: Various Classrooms\n\nCome grow in God's Word with fellow believers!`,
            time: "10:30 AM",
            location: "Various Classrooms",
          },
          youth: {
            title: "Youth Group Meeting",
            summary: "Fellowship and fun for teens and young adults.",
            content: `Our youth group meets every Saturday for games, discussion, and growing together in faith.\n\nAges: 13-25\nActivities include:\n- Group games\n- Bible discussion\n- Prayer time\n- Snacks and fellowship\n\nCome join us for an evening of fun and spiritual growth!`,
            time: "6:00 PM",
            location: "Youth Center",
          },
          prayerNight: {
            title: "Prayer Night",
            summary: "Corporate prayer and intercession.",
            content: `Join us for our monthly prayer night where we lift up the needs of our church, community, and world.\n\nWe will pray for:\n- Our church family\n- Community needs\n- Mission work\n- Personal requests\n\nAll are welcome to join us in prayer.`,
            time: "7:00 PM",
            location: "Main Sanctuary",
          },
          baptism: {
            title: "Baptism Service",
            summary: "Public baptism and celebration of new life in Christ.",
            content: `Join us as we celebrate baptisms of new believers publicly professing their faith in Christ.\n\nThis is a special time of worship and celebration as we witness these powerful testimonies of God's grace.\n\nAfter the service, we'll have a fellowship celebration with refreshments.`,
            time: "After Service",
            location: "Main Sanctuary",
          },
        },
      },
      kh: {
        hero: {
          eyebrow: "ប្រតិទិនសេវាកម្ម",
          title: "ប្រតិទិនព្រះវិហារ",
          subtitle:
            "មើលសកម្មភាពសប្ដាហ៍ ព្រឹត្តិការណ៍ជាក់លាក់ និងការជួបជុំគ្រួសារទាំងអស់នៅកន្លែងតែមួយ។",
          chips: [
            "កាលវិភាគប្រចាំខែ",
            "ថ្វាយបង្គំ និងអធិស្ឋាន",
            "យុវជន និងគ្រួសារ",
          ],
          checklist: [
            "រុករកតាមខែ ហើយចុចលើថ្ងៃដែលបន្លិច។",
            "ទទួលម៉ោង ទីតាំង និងព័ត៌មានលម្អិតពេញលេញ។",
            "ស្វែងរកសេវាកម្ម ថ្នាក់រៀន និងយប់អធិស្ឋានបានយ៉ាងរហ័ស។",
          ],
        },
        highlights: {
          title: "គន្លឹះខែនេះ",
          footer: "ព្រឹត្តិការណ៍ជាក់ស្ដែង · ព័ត៌មានទាន់ហេតុការណ៍",
          status: "សកម្ម",
          items: [
            {
              title: "សេចក្ដីថ្វាយបង្គំព្រះអាទិត្យ",
              detail: "រៀងរាល់អាទិត្យ · ម៉ោង ៩:០០ ព្រឹក",
            },
            { title: "ក្រុមយុវជន", detail: "រៀងរាល់សៅរ៍ · ម៉ោង ៦:០០ ល្ងាច" },
            {
              title: "យប់អធិស្ឋាន",
              detail: "សៅរ៍ដំបូងប្រចាំខែ · ម៉ោង ៧:០០ ល្ងាច",
            },
          ],
        },
        how: {
          eyebrow: "របៀបប្រើប្រាស់",
          title: "ប្រើប្រតិទិនយ៉ាងឆាប់រហ័ស",
          helper: "ជ្រើសថ្ងៃដែលបន្លិចដើម្បីមើលព្រឹត្តិការណ៍ និងព័ត៌មានលម្អិត។",
          steps: [
            {
              title: "រុករកតាមខែ",
              body: 'ប្រើប៊ូតុងខែមុន/ខែក្រោយ ឬ "ថ្ងៃនេះ" ដើម្បីត្រឡប់ទៅខែបច្ចុប្បន្ន។',
            },
            {
              title: "ជ្រើសថ្ងៃដែលបន្លិច",
              body: "ថ្ងៃដែលបន្លិចបង្ហាញចំនួនព្រឹត្តិការណ៍ក្នុងថ្ងៃនោះ។",
            },
            {
              title: "បើកព្រឹត្តិការណ៍",
              body: "មើលម៉ោង ទីតាំង និងការពិពណ៌នា ហើយចែករំលែកព័ត៌មានជាមួយមិត្តភក្តិ។",
            },
          ],
        },
        calendar: {
          helper: "ជ្រើសថ្ងៃដែលបន្លិចដើម្បីមើលព្រឹត្តិការណ៍ និងព័ត៌មានលម្អិត។",
          weekdays: [
            "អាទិត្យ",
            "ចន្ទ",
            "អង្គារ",
            "ពុធ",
            "ព្រហស្បតិ៍",
            "សុក្រ",
            "សៅរ៍",
          ],
          eventCountLabel: "ព្រឹត្តិការណ៍",
          moreLabel: "ទៀត",
          selectedPrefix: "ព្រឹត្តិការណ៍នៅ",
          today: "ថ្ងៃនេះ",
          prevAria: "ខែមុន",
          nextAria: "ខែក្រោយ",
        },
        loading: "កំពុងផ្ទុកប្រតិទិន...",
        error: "មិនអាចផ្ទុកប្រតិទិនបានទេ",
        months: [
          "មករា",
          "កុម្ភៈ",
          "មីនា",
          "មេសា",
          "ឧសភា",
          "មិថុនា",
          "កក្កដា",
          "សីហា",
          "កញ្ញា",
          "តុលា",
          "វិច្ឆិកា",
          "ធ្នូ",
        ],
        events: {
          sunday: {
            title: "សេចក្ដីថ្វាយបង្គំព្រះអាទិត្យ",
            summary:
              "ចូលរួមរៀងរាល់អាទិត្យសម្រាប់ការថ្វាយបង្គំ អធិស្ឋាន និងសមាគម។",
            content: `យើងជួបគ្នារៀងរាល់ព្រឹកថ្ងៃអាទិត្យសម្រាប់ការថ្វាយបង្គំ និងស្តាប់បដិប្បធម៌។\n\nកាលវិភាគសេវាកម្ម៖\n- សរសើរ និងថ្វាយបង្គំ៖ ៩:០០ ព្រឹក\n- អធិស្ឋាន៖ ៩:៣០ ព្រឹក\n- សារៈ៖ ១០:០០ ព្រឹក\n- សមាគម៖ ១១:៣០ ព្រឹក\n\nអ្នករាល់គ្នាត្រូវបានស្វាគមន៍! យើងរីករាយក្នុងការថ្វាយបង្គំជាមួយអ្នក។`,
            time: "9:00 AM",
            location: "សាលាថ្វាយបង្គំ",
          },
          sundaySchool: {
            title: "ថ្នាក់រៀនព្រះគម្ពីរ",
            summary: "ថ្នាក់សិក្សាព្រះគម្ពីរសម្រាប់គ្រប់វ័យ។",
            content: `ចូលរួមថ្នាក់រៀនព្រះគម្ពីរសម្រាប់កុមារ យុវជន និងមនុស្សពេញវ័យ។\n\nថ្នាក់រៀនទាំងអស់ចាប់ពី ១០:៣០ ព្រឹក ដល់ ១១:៣០ ព្រឹក\nទីតាំង៖ បន្ទប់សិក្សាផ្សេងៗ\n\nមករីកចម្រើនក្នុងព្រះបន្ទូលជាមួយបងប្អូនជំនឿ!`,
            time: "10:30 AM",
            location: "បន្ទប់សិក្សាផ្សេងៗ",
          },
          youth: {
            title: "ជំនុំក្រុមយុវជន",
            summary: "សមាគម និងកម្សាន្តសម្រាប់យុវជន។",
            content: `ក្រុមយុវជនជួបគ្នារៀងរាល់ថ្ងៃសៅរ៍សម្រាប់ហ្គេម ការពិភាក្សា និងលូតលាស់ជំនឿរួមគ្នា។\n\nអាយុ៖ ១៣-២៥ ឆ្នាំ\nសកម្មភាពរួមមាន៖\n- ហ្គេមជាក្រុម\n- ពិភាក្សាព្រះគម្ពីរ\n- ពេលវេលាអធិស្ឋាន\n- អាហារសម្រន់ និងសមាគម\n\nសូមចូលរួមយប់កម្សាន្ត និងលូតលាស់វិញ្ញាណជាមួយពួកយើង!`,
            time: "6:00 PM",
            location: "មណ្ឌលយុវជន",
          },
          prayerNight: {
            title: "យប់អធិស្ឋាន",
            summary: "អធិស្ឋានរួម និងអន្តរាស្ត្រ។",
            content: `ចូលរួមយប់អធិស្ឋានប្រចាំខែ ដើម្បីអធិស្ឋានសម្រាប់តម្រូវការរបស់ព្រះវិហារ សហគមន៍ និងពិភពលោក។\n\nយើងនឹងអធិស្ឋានសម្រាប់៖\n- គ្រួសារព្រះវិហារ\n- តម្រូវការសហគមន៍\n- ការសម្រេចបេសកកម្ម\n- សំណូមពរផ្ទាល់ខ្លួន\n\nសូមស្វាគមន៍បងប្អូនចូលរួមអធិស្ឋានជាមួយយើង។`,
            time: "7:00 PM",
            location: "សាលាថ្វាយបង្គំ",
          },
          baptism: {
            title: "ពិធីបុណ្យជ្រមុជទឹក",
            summary:
              "ពិធីជ្រមុជទឹកជាសាធារណៈ និងអបអរសាទរជីវិតថ្មីក្នុងព្រះគ្រីស្ទ។",
            content: `ចូលរួមជាមួយយើងក្នុងការអបអរសាទរពិធីជ្រមុជទឹករបស់អ្នកជំនឿថ្មីដែលប្រកាសជំនឿជាសាធារណៈ។\n\nនេះជាពេលពិសេសនៃការថ្វាយបង្គំ និងអបអរសាទរ ខណៈដែលយើងជាភស្តុតាងដល់សក្ដានុពលនៃព្រះគុណ។\n\nបន្ទាប់ពីសេវាកម្ម យើងនឹងមានសមាគមជាមួយអាហារសម្រន់។`,
            time: "ក្រោយសេវាកម្ម",
            location: "សាលាថ្វាយបង្គំ",
          },
        },
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    const hero = computed(() => content.value.hero);
    const how = computed(() => content.value.how);
    const highlights = computed(() => content.value.highlights);
    const calendarCopy = computed(() => content.value.calendar);
    const weekdayLabels = computed(() => calendarCopy.value.weekdays);
    const loadingText = computed(() => content.value.loading);

    const formatFullDate = (date) => {
      if (!date) return "";
      const dateObj = date instanceof Date ? date : new Date(date);
      if (isNaN(dateObj.getTime())) return "";
      const months = content.value.months;
      // Use UTC methods for consistent date display
      return `${months[dateObj.getUTCMonth()]} ${dateObj.getUTCDate()}, ${dateObj.getUTCFullYear()}`;
    };

    const formatTimeRange = (start, end) => {
      if (!start) return "";
      // Format time using UTC to ensure consistency
      const formatUTCTime = (date) => {
        const hours = String(date.getUTCHours()).padStart(2, '0');
        const minutes = String(date.getUTCMinutes()).padStart(2, '0');
        return `${hours}:${minutes}`;
      };
      const startStr = formatUTCTime(start);
      const endStr = end ? formatUTCTime(end) : "";
      return endStr && endStr !== startStr
        ? `${startStr} - ${endStr}`
        : startStr;
    };

    const normalizeEvents = (apiEvents) =>
      apiEvents.map((event) => {
        // CRITICAL: Extract date keys from raw API strings BEFORE any Date parsing
        // Get the raw string values from API response (these are ISO 8601 UTC strings)
        const startDateStr = typeof event.startDate === 'string' 
          ? event.startDate 
          : (event.startDate ? event.startDate.toISOString() : (event.startTime || event.date || ''))
        const endDateStr = typeof event.endDate === 'string'
          ? event.endDate
          : (event.endDate ? event.endDate.toISOString() : (event.endTime || startDateStr))
        
        // Parse as Date objects ONLY for time operations (display, comparison)
        const start = new Date(startDateStr);
        const end = endDateStr ? new Date(endDateStr) : start;

        return {
          ...event,
          startTime: start,
          endTime: end,
          // CRITICAL: Keep original raw strings for date key extraction
          _startTimeStr: startDateStr,
          _endTimeStr: endDateStr,
          recurrence: (event.recurrence || "none").toLowerCase(),
          recurrenceEndsAt: event.recurrenceEndsAt
            ? new Date(event.recurrenceEndsAt)
            : null,
          timeRange: formatTimeRange(start, end),
          category: event.category || "",
          color: event.color || "",
          content: event.details || event.description || event.summary || event.content || "",
        };
      });

    const addInterval = (date, recurrence) => {
      const next = new Date(date);
      // Use UTC methods to maintain UTC consistency
      switch (recurrence) {
        case "daily":
          next.setUTCDate(next.getUTCDate() + 1);
          break;
        case "weekly":
          next.setUTCDate(next.getUTCDate() + 7);
          break;
        case "monthly":
          next.setUTCMonth(next.getUTCMonth() + 1);
          break;
        case "yearly":
          next.setUTCFullYear(next.getUTCFullYear() + 1);
          break;
        default:
          break;
      }
      return next;
    };

    const expandEventOccurrences = (event, rangeStart, rangeEnd) => {
      const occurrences = [];
      const recurrence = event.recurrence || "none";
      const recurrenceEnd = event.recurrenceEndsAt
        ? new Date(event.recurrenceEndsAt)
        : null;
      const maxIterations = 500;

      // event.startTime and event.endTime are already Date objects from normalizeEvents
      let currentStart = new Date(event.startTime);
      let currentEnd = new Date(event.endTime);
      let iterations = 0;

      const pushIfInRange = (start, end) => {
        // Compare dates directly (Date objects compare by timestamp)
        if (end < rangeStart || start > rangeEnd) {
          return;
        }

        // Extract date key from UTC ISO string for consistent matching
        const occurrenceDateKey = start.toISOString().split('T')[0];
        
        occurrences.push({
          ...event,
          occurrenceId: `${event.id}-${start.toISOString()}`,
          startTime: start,
          endTime: end,
          timeRange: formatTimeRange(start, end),
          // Store the date key for this occurrence
          _startTimeStr: start.toISOString(),
          _endTimeStr: end.toISOString(),
        });
      };

      while (iterations < maxIterations && currentStart <= rangeEnd) {
        if (!recurrenceEnd || currentStart <= recurrenceEnd) {
          pushIfInRange(new Date(currentStart), new Date(currentEnd));
        }

        if (recurrence === "none") {
          break;
        }

        currentStart = addInterval(currentStart, recurrence);
        currentEnd = addInterval(currentEnd, recurrence);
        iterations++;
      }

      return occurrences;
    };

    const fetchEvents = async () => {
      loading.value = true;

      try {
        const response = await eventService.getEvents(false);
        const apiEvents = response.events || [];
        console.log('Fetched events from API:', apiEvents);
        events.value = normalizeEvents(apiEvents);
        console.log('Normalized events:', events.value);
        loading.value = false;
      } catch (err) {
        console.error("Error loading events:", err);
        error.value = err.message || content.value.error;
        loading.value = false;
      }
    };

    watch(selectedLanguage, () => {
      error.value = "";
      fetchEvents();
    });

    const calendarRange = computed(() => {
      const year = currentYear.value;
      const month = currentMonth.value;

      // Build calendar range using pure UTC date math - no Date object manipulation
      // Calculate first and last day of month
      const firstDayOfMonth = { year, month, day: 1 };
      const lastDayOfMonthDate = new Date(Date.UTC(year, month + 1, 0));
      const lastDayOfMonth = { year, month, day: lastDayOfMonthDate.getUTCDate() };
      
      // Calculate what day of week the first day falls on (0=Sunday, 6=Saturday)
      const firstDayDate = new Date(Date.UTC(year, month, 1));
      const firstDayOfWeek = firstDayDate.getUTCDay();
      
      // Start date is first day of month minus days to get to Sunday
      let startYear = year;
      let startMonth = month;
      let startDay = 1 - firstDayOfWeek;
      
      // Handle case where startDay goes negative (previous month)
      if (startDay <= 0) {
        startMonth--;
        if (startMonth < 0) {
          startMonth = 11;
          startYear--;
        }
        const daysInPrevMonth = new Date(Date.UTC(startYear, startMonth + 1, 0)).getUTCDate();
        startDay = daysInPrevMonth + startDay;
      }
      
      // Calculate what day of week the last day falls on
      const lastDayDate = new Date(Date.UTC(year, month, lastDayOfMonth.day));
      const lastDayOfWeek = lastDayDate.getUTCDay();
      
      // End date is last day of month plus days to get to Saturday
      let endYear = year;
      let endMonth = month;
      let endDay = lastDayOfMonth.day + (6 - lastDayOfWeek);
      
      // Handle case where endDay exceeds days in month (next month)
      const daysInEndMonth = new Date(Date.UTC(endYear, endMonth + 1, 0)).getUTCDate();
      if (endDay > daysInEndMonth) {
        endDay = endDay - daysInEndMonth;
        endMonth++;
        if (endMonth > 11) {
          endMonth = 0;
          endYear++;
        }
      }
      
      return { 
        start: { year: startYear, month: startMonth, day: startDay },
        end: { year: endYear, month: endMonth, day: endDay }
      };
    });

    const expandedEvents = computed(() => {
      const { start, end } = calendarRange.value;
      // Convert calendar range to Date objects for comparison
      const startDate = new Date(Date.UTC(start.year, start.month, start.day));
      const endDate = new Date(Date.UTC(end.year, end.month, end.day));
      // Set endDate to end of day
      endDate.setUTCHours(23, 59, 59, 999);
      
      const list = [];

      events.value.forEach((event) => {
        list.push(...expandEventOccurrences(event, startDate, endDate));
      });

      return list;
    });

    const eventsByDate = computed(() => {
      const map = {};

      expandedEvents.value.forEach((event) => {
        // CRITICAL: Extract date key ONLY from the raw ISO string
        // This is the ONLY way to guarantee no timezone conversion
        if (!event._startTimeStr) {
          console.warn('Event missing _startTimeStr:', event)
          return
        }
        
        // Extract date directly from ISO string: "2026-01-30T09:00:00Z" -> "2026-01-30"
        // Split on 'T' and take first part - this is the UTC date as stored
        let dateKey = event._startTimeStr.split('T')[0]
        
        // Handle case where there might be a space instead of T (some API formats)
        if (!dateKey && event._startTimeStr.includes(' ')) {
          dateKey = event._startTimeStr.split(' ')[0]
        }
        
        // Validate the date key format (should be YYYY-MM-DD)
        if (!/^\d{4}-\d{2}-\d{2}$/.test(dateKey)) {
          console.warn('Invalid date key format:', dateKey, 'from:', event._startTimeStr, 'event:', event)
          return
        }
        
        if (!map[dateKey]) {
          map[dateKey] = [];
        }
        map[dateKey].push(event);
      });

      Object.values(map).forEach((list) => {
        list.sort((a, b) => a.startTime - b.startTime);
      });

      // Debug: Log events by date (only first time)
      if (Object.keys(map).length > 0 && !window._calendarEventsLogged) {
        console.log('Calendar events by date:', map);
        console.log('Total events:', expandedEvents.value.length);
        window._calendarEventsLogged = true;
      }

      return map;
    });

    const calendarDays = computed(() => {
      const { start, end } = calendarRange.value;
      const month = currentMonth.value;

      const days = [];
      
      // Work with pure year/month/day numbers - no Date object manipulation
      let year = start.year;
      let monthNum = start.month;
      let day = start.day;
      
      const endYear = end.year;
      const endMonth = end.month;
      const endDay = end.day;
      
      // Loop through days using pure date math
      while (true) {
        // Build UTC date key directly from year/month/day - exactly like event date keys
        const dateKey = `${year}-${String(monthNum + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
        
        // Check if we've reached the end
        if (year > endYear || 
            (year === endYear && monthNum > endMonth) ||
            (year === endYear && monthNum === endMonth && day > endDay)) {
          break;
        }
        
        const dayEvents = eventsByDate.value[dateKey] || [];

        // Compare dates for today using pure date comparison
        const today = new Date();
        const todayYear = today.getUTCFullYear();
        const todayMonth = today.getUTCMonth();
        const todayDay = today.getUTCDate();
        
        const isToday = year === todayYear && monthNum === todayMonth && day === todayDay;

        days.push({
          date: dateKey,
          day: day, // Use the day we're tracking
          isCurrentMonth: monthNum === currentMonth.value,
          isToday: isToday,
          hasEvents: dayEvents.length > 0,
          events: dayEvents,
        });

        // Increment to next day using pure date math
        const daysInMonth = new Date(Date.UTC(year, monthNum + 1, 0)).getUTCDate();
        day++;
        if (day > daysInMonth) {
          day = 1;
          monthNum++;
          if (monthNum > 11) {
            monthNum = 0;
            year++;
          }
        }
      }

      return days;
    });

    const selectedDateEvents = computed(() => {
      if (!selectedDay.value) return [];

      // Build UTC date key directly from UTC components
      const year = selectedDay.value.getUTCFullYear();
      const month = String(selectedDay.value.getUTCMonth() + 1).padStart(2, '0');
      const day = String(selectedDay.value.getUTCDate()).padStart(2, '0');
      const dateKey = `${year}-${month}-${day}`;
      
      return eventsByDate.value[dateKey] || [];
    });

    const currentMonthName = computed(
      () => content.value.months[currentMonth.value],
    );

    const selectDate = (day) => {
      if (!day.isCurrentMonth) return;
      // day.date is a UTC date string (YYYY-MM-DD), create UTC date
      const [year, month, date] = day.date.split("-").map(Number);
      selectedDay.value = new Date(Date.UTC(year, month - 1, date));
      
      // If the day has events, open the modal with the first event
      if (day.events && day.events.length > 0) {
        openModal(day.events[0]);
      }
    };

    const previousMonth = () => {
      if (currentMonth.value === 0) {
        currentMonth.value = 11;
        currentYear.value--;
      } else {
        currentMonth.value--;
      }
      selectedDay.value = null;
    };

    const nextMonth = () => {
      if (currentMonth.value === 11) {
        currentMonth.value = 0;
        currentYear.value++;
      } else {
        currentMonth.value++;
      }
      selectedDay.value = null;
    };

    const resetToCurrentMonth = () => {
      const now = new Date();
      currentMonth.value = now.getUTCMonth();
      currentYear.value = now.getUTCFullYear();
      selectedDay.value = null;
    };

    const openModal = (event) => {
      selectedEvent.value = event;
    };

    const closeModal = () => {
      selectedEvent.value = null;
    };

    onMounted(() => {
      fetchEvents();
    });

    return {
      loading,
      loadingText,
      error,
      hero,
      how,
      highlights,
      calendarCopy,
      weekdayLabels,
      selectedEvent,
      currentMonth,
      currentYear,
      selectedDay,
      selectedDateEvents,
      currentMonthName,
      calendarDays,
      formatFullDate,
      selectDate,
      previousMonth,
      nextMonth,
      resetToCurrentMonth,
      openModal,
      closeModal,
    };
  },
};
</script>
