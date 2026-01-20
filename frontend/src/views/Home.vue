<template>
  <div>
    <!-- Hero Section -->
    <div
      class="relative min-h-screen flex items-center overflow-hidden pt-24 md:pt-28"
    >
      <div class="absolute inset-0 overflow-hidden" aria-hidden="true">
        <div
          class="absolute inset-0 bg-cover bg-center scale-105 saturate-110 brightness-105 transition-transform duration-[4000ms]"
          :style="heroBackgroundStyle"
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
        class="relative z-10 w-full max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 text-center"
        aria-live="polite"
      >
        <p class="text-sm uppercase tracking-[0.4em] text-white/70 mb-4">
          {{ activeSlide.eyebrow }}
        </p>
        <h1
          class="text-4xl md:text-5xl lg:text-6xl font-bold text-white leading-tight mb-6"
        >
          {{ activeSlide.title }}
        </h1>
        <p
          class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light mb-10"
        >
          {{ activeSlide.description }}
        </p>
        <div class="flex flex-wrap justify-center gap-4">
          <router-link
            v-if="activeSlide.primaryCta"
            :to="activeSlide.primaryCta.route"
            class="inline-flex items-center px-6 py-3 bg-brand-orange text-white font-semibold rounded-md hover:bg-brand-orange/90 transition-colors"
          >
            {{ activeSlide.primaryCta.label }}
          </router-link>
          <router-link
            v-if="activeSlide.secondaryCta"
            :to="activeSlide.secondaryCta.route"
            class="inline-flex items-center px-6 py-3 bg-white/10 text-white font-semibold rounded-md border border-white/20 hover:bg-white/20 transition-colors backdrop-blur-sm"
          >
            {{ activeSlide.secondaryCta.label }}
          </router-link>
        </div>
      </div>

      <button
        type="button"
        class="absolute left-4 sm:left-6 lg:left-10 top-1/2 -translate-y-1/2 inline-flex items-center justify-center p-3 sm:p-4 rounded-full border border-white/40 text-white bg-white/10 hover:bg-white/20 transition-all focus:outline-none focus:ring-2 focus:ring-white/80"
        aria-label="Previous banner"
        @click="prevSlide"
      >
        <svg
          class="w-5 h-5 sm:w-6 sm:h-6"
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
        type="button"
        class="absolute right-4 sm:right-6 lg:right-10 top-1/2 -translate-y-1/2 inline-flex items-center justify-center p-3 sm:p-4 rounded-full border border-white/40 text-white bg-white/10 hover:bg-white/20 transition-all focus:outline-none focus:ring-2 focus:ring-white/80"
        aria-label="Next banner"
        @click="nextSlide"
      >
        <svg
          class="w-5 h-5 sm:w-6 sm:h-6"
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

      <div
        class="absolute inset-x-0 bottom-8 px-6 sm:px-10 z-20 flex items-center justify-center"
      >
        <div class="flex items-center gap-3" aria-label="Select hero banner">
          <button
            v-for="(slide, index) in heroSlides"
            :key="slide.id"
            type="button"
            :aria-label="`Show banner ${index + 1}`"
            :aria-current="currentSlideIndex === index ? 'true' : 'false'"
            class="h-2.5 rounded-full transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-white/80 focus:ring-offset-2 focus:ring-offset-white/10"
            :class="
              currentSlideIndex === index
                ? 'w-9 bg-white'
                : 'w-2.5 bg-white/40 hover:bg-white/70'
            "
            @click="goToSlide(index)"
          ></button>
        </div>
      </div>
    </div>

    <!-- Regular Study Times -->
    <section class="py-20 bg-gradient-to-b from-white to-gray-50">
      <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400 mb-3">
            {{ weekly.eyebrow }}
          </p>
          <h2 class="text-4xl font-bold text-gray-900 mb-4">
            {{ weekly.title }}
          </h2>
          <p class="text-lg text-gray-600 max-w-3xl mx-auto">
            {{ weekly.description }}
          </p>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div
            class="relative overflow-hidden rounded-3xl border border-gray-100 bg-white shadow-2xl"
          >
            <div
              class="absolute inset-0 opacity-60 pointer-events-none"
              :style="{ background: weeklyGatherings[0].backgroundStyle }"
            ></div>
            <div
              class="relative flex flex-col md:flex-row items-center md:items-center gap-8 p-8"
            >
              <div
                class="flex flex-col items-center justify-center text-center rounded-2xl px-6 py-7 w-full md:w-40 shrink-0"
                :class="
                  weeklyGatherings[0].accent === 'orange'
                    ? 'bg-brand-orange/10 text-brand-orange'
                    : 'bg-brand-blue/10 text-brand-blue'
                "
              >
                <span
                  class="text-xs uppercase tracking-[0.4em]"
                  :class="
                    weeklyGatherings[0].accent === 'orange'
                      ? 'text-brand-orange/80'
                      : 'text-brand-blue/70'
                  "
                >
                  {{ weeklyGatherings[0].dayLabel }}
                </span>
                <span
                  class="text-4xl md:text-5xl font-bold text-gray-900 mt-2 leading-none"
                  >{{ weeklyGatherings[0].time }}</span
                >
                <span class="text-sm tracking-[0.3em] text-gray-500 mt-1">{{
                  weeklyGatherings[0].meridiem
                }}</span>
              </div>
              <div class="flex-1 text-gray-700">
                <div
                  class="inline-flex items-center px-3 py-1 rounded-full uppercase tracking-[0.35em] text-xs font-semibold mb-4"
                  :class="
                    weeklyGatherings[0].accent === 'orange'
                      ? 'bg-brand-orange/10 text-brand-orange'
                      : 'bg-brand-blue/10 text-brand-blue'
                  "
                >
                  {{ weeklyGatherings[0].badge }}
                </div>
                <h3 class="text-3xl font-semibold text-gray-900 mb-3">
                  {{ weeklyGatherings[0].title }}
                </h3>
                <p class="text-base leading-relaxed text-gray-600 mb-5">
                  {{ weeklyGatherings[0].body }}
                </p>
                <router-link
                  :to="weeklyGatherings[0].ctaRoute"
                  class="inline-flex items-center font-semibold transition-colors"
                  :class="
                    weeklyGatherings[0].accent === 'orange'
                      ? 'text-brand-orange hover:text-brand-blue'
                      : 'text-brand-blue hover:text-brand-orange'
                  "
                >
                  {{ weeklyGatherings[0].ctaLabel }}
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
            </div>
          </div>

          <div
            class="relative overflow-hidden rounded-3xl border border-gray-100 bg-white shadow-2xl"
          >
            <div
              class="absolute inset-0 opacity-60 pointer-events-none"
              :style="{ background: weeklyGatherings[1].backgroundStyle }"
            ></div>
            <div
              class="relative flex flex-col md:flex-row items-center md:items-center gap-8 p-8"
            >
              <div
                class="flex flex-col items-center justify-center text-center rounded-2xl px-6 py-7 w-full md:w-40 shrink-0"
                :class="
                  weeklyGatherings[1].accent === 'orange'
                    ? 'bg-brand-orange/10 text-brand-orange'
                    : 'bg-brand-blue/10 text-brand-blue'
                "
              >
                <span
                  class="text-xs uppercase tracking-[0.4em]"
                  :class="
                    weeklyGatherings[1].accent === 'orange'
                      ? 'text-brand-orange/80'
                      : 'text-brand-blue/70'
                  "
                >
                  {{ weeklyGatherings[1].dayLabel }}
                </span>
                <span
                  class="text-4xl md:text-5xl font-bold text-gray-900 mt-2 leading-none"
                  >{{ weeklyGatherings[1].time }}</span
                >
                <span class="text-sm tracking-[0.3em] text-gray-500 mt-1">{{
                  weeklyGatherings[1].meridiem
                }}</span>
              </div>
              <div class="flex-1 text-gray-700">
                <div
                  class="inline-flex items-center px-3 py-1 rounded-full uppercase tracking-[0.35em] text-xs font-semibold mb-4"
                  :class="
                    weeklyGatherings[1].accent === 'orange'
                      ? 'bg-brand-orange/10 text-brand-orange'
                      : 'bg-brand-blue/10 text-brand-blue'
                  "
                >
                  {{ weeklyGatherings[1].badge }}
                </div>
                <h3 class="text-3xl font-semibold text-gray-900 mb-3">
                  {{ weeklyGatherings[1].title }}
                </h3>
                <p class="text-base leading-relaxed text-gray-600 mb-5">
                  {{ weeklyGatherings[1].body }}
                </p>
                <router-link
                  :to="weeklyGatherings[1].ctaRoute"
                  class="inline-flex items-center font-semibold transition-colors"
                  :class="
                    weeklyGatherings[1].accent === 'orange'
                      ? 'text-brand-orange hover:text-brand-blue'
                      : 'text-brand-blue hover:text-brand-orange'
                  "
                >
                  {{ weeklyGatherings[1].ctaLabel }}
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
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Upcoming Section -->
    <div class="relative py-24 bg-brand-blue overflow-hidden">
      <div class="absolute inset-0 opacity-80" aria-hidden="true">
        <div
          class="absolute inset-y-0 left-0 w-1/2 bg-gradient-to-r from-black/40 via-transparent to-transparent"
        ></div>
        <div
          class="absolute inset-y-0 right-0 w-2/3 bg-gradient-to-l from-brand-orange/20 via-transparent to-transparent"
        ></div>
        <div
          class="absolute inset-x-0 top-0 h-1/2 bg-gradient-to-b from-white/15 via-transparent to-transparent"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-56 bg-gradient-to-t from-black/30 via-brand-blue/20 to-transparent"
        ></div>
        <div
          class="absolute inset-0 opacity-35 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 25% 20%,
                rgba(255, 255, 255, 0.25),
                transparent 40%
              ),
              radial-gradient(
                circle at 80% 0%,
                rgba(255, 255, 255, 0.15),
                transparent 35%
              ),
              radial-gradient(
                circle at 60% 85%,
                rgba(255, 255, 255, 0.2),
                transparent 40%
              );
          "
        ></div>
      </div>
      <div class="relative">
        <div class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 text-white">
          <div class="text-center mb-16">
            <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">
              {{ upcoming.eyebrow }}
            </p>
            <h2 class="text-4xl md:text-5xl font-bold mb-4">
              {{ upcoming.title }}
            </h2>
            <p class="text-white/80 max-w-2xl mx-auto">
              {{ upcoming.description }}
            </p>
          </div>

          <div class="max-w-2xl mx-auto space-y-6">
            <div
              class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]"
            >
              <div
                class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8"
              >
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2"
                  >{{ upcomingItems[0].month }}</span
                >
                <span class="text-5xl font-bold leading-none text-brand-blue">{{
                  upcomingItems[0].day
                }}</span>
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3"
                  >{{ upcomingItems[0].weekday }}</span
                >
              </div>
              <div
                class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8"
              >
                <div class="flex-1">
                  <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">
                    {{ upcomingItems[0].title }}
                  </h3>
                  <p class="text-white/80 text-base leading-relaxed">
                    {{ upcomingItems[0].description }}
                  </p>
                </div>
                <router-link
                  :to="upcomingItems[0].route"
                  class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
                >
                  {{ upcoming.ctaLabel }}
                  <svg
                    class="w-4 h-4"
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
            </div>

            <div
              class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]"
            >
              <div
                class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8"
              >
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2"
                  >{{ upcomingItems[1].month }}</span
                >
                <span class="text-5xl font-bold leading-none text-brand-blue">{{
                  upcomingItems[1].day
                }}</span>
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3"
                  >{{ upcomingItems[1].weekday }}</span
                >
              </div>
              <div
                class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8"
              >
                <div class="flex-1">
                  <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">
                    {{ upcomingItems[1].title }}
                  </h3>
                  <p class="text-white/80 text-base leading-relaxed">
                    {{ upcomingItems[1].description }}
                  </p>
                </div>
                <router-link
                  :to="upcomingItems[1].route"
                  class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
                >
                  {{ upcoming.ctaLabel }}
                  <svg
                    class="w-4 h-4"
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
            </div>

            <div
              class="bg-white/10 rounded-2xl border border-white/15 px-6 py-8 backdrop-blur-md shadow-lg flex flex-col gap-6 sm:flex-row sm:items-stretch min-h-[200px]"
            >
              <div
                class="flex flex-col items-center justify-center text-center w-full sm:w-40 sm:min-w-[10rem] bg-white text-brand-blue rounded-2xl border border-white/30 shadow-2xl px-6 py-8"
              >
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mb-2"
                  >{{ upcomingItems[2].month }}</span
                >
                <span class="text-5xl font-bold leading-none text-brand-blue">{{
                  upcomingItems[2].day
                }}</span>
                <span
                  class="text-xs uppercase tracking-[0.4em] text-brand-blue/60 mt-3"
                  >{{ upcomingItems[2].weekday }}</span
                >
              </div>
              <div
                class="flex-1 flex flex-col gap-4 sm:flex-row sm:items-center sm:gap-8"
              >
                <div class="flex-1">
                  <h3 class="text-xl md:text-2xl font-semibold text-white mb-2">
                    {{ upcomingItems[2].title }}
                  </h3>
                  <p class="text-white/80 text-base leading-relaxed">
                    {{ upcomingItems[2].description }}
                  </p>
                </div>
                <router-link
                  :to="upcomingItems[2].route"
                  class="inline-flex items-center gap-2 px-5 py-2.5 bg-white text-brand-blue font-semibold rounded-full tracking-wide text-sm hover:bg-brand-orange hover:text-white transition-colors whitespace-nowrap self-start sm:self-auto"
                >
                  {{ upcoming.ctaLabel }}
                  <svg
                    class="w-4 h-4"
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
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Latest Teachings Section -->
    <div
      class="relative py-24 bg-gradient-to-br from-[#0B111E] via-[#112437] to-[#1C4053] text-white overflow-hidden"
    >
      <div class="absolute inset-0 opacity-70" aria-hidden="true">
        <div
          class="absolute inset-0"
          style="
            background-image:
              radial-gradient(
                circle at 20% 15%,
                rgba(228, 157, 64, 0.25),
                transparent 45%
              ),
              radial-gradient(
                circle at 80% 10%,
                rgba(255, 255, 255, 0.12),
                transparent 40%
              ),
              radial-gradient(
                circle at 60% 85%,
                rgba(34, 112, 130, 0.3),
                transparent 45%
              );
          "
        ></div>
        <div
          class="absolute inset-0 mix-blend-soft-light opacity-40"
          style="
            background-image:
              linear-gradient(
                120deg,
                rgba(255, 255, 255, 0.15) 0%,
                rgba(255, 255, 255, 0) 45%
              ),
              linear-gradient(
                300deg,
                rgba(34, 112, 130, 0.25) 0%,
                rgba(255, 255, 255, 0) 55%
              );
          "
        ></div>
        <div
          class="absolute inset-0 opacity-20"
          style="
            background-image:
              linear-gradient(
                to right,
                rgba(255, 255, 255, 0.06) 1px,
                transparent 1px
              ),
              linear-gradient(
                to bottom,
                rgba(255, 255, 255, 0.06) 1px,
                transparent 1px
              );
            background-size: 80px 80px;
          "
        ></div>
      </div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-16">
          <p class="text-sm uppercase tracking-[0.3em] text-white/70 mb-3">
            {{ teachings.eyebrow }}
          </p>
          <h2 class="text-4xl md:text-5xl font-bold text-white mb-4">
            {{ teachings.title }}
          </h2>
          <p class="text-lg text-white/80 max-w-2xl mx-auto font-light">
            {{ teachings.description }}
          </p>
        </div>

        <!-- Search Bar -->
        <div class="max-w-2xl mx-auto mb-12">
          <div class="flex gap-3">
            <input
              v-model="videoSearchQuery"
              type="text"
              :placeholder="teachings.searchPlaceholder || 'Search videos...'"
              class="flex-1 px-4 py-3 bg-white/10 border border-white/20 rounded-full text-white placeholder-white/50 focus:outline-none focus:ring-2 focus:ring-white/50 focus:border-white/40 backdrop-blur-sm"
              @keyup.enter="performVideoSearch"
            />
            <button
              type="button"
              @click="performVideoSearch"
              class="px-6 py-3 bg-brand-orange text-white font-semibold rounded-full hover:bg-brand-orange/90 transition-colors focus:outline-none focus:ring-2 focus:ring-brand-orange/50"
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
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </button>
            <button
              v-if="videoSearchQuery"
              type="button"
              @click="clearVideoSearch"
              class="px-6 py-3 bg-white/10 border border-white/20 text-white font-semibold rounded-full hover:bg-white/20 transition-colors focus:outline-none focus:ring-2 focus:ring-white/50"
            >
              Clear
            </button>
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="teaching in displayedTeachings"
            :key="teaching.id"
            class="bg-white/10 border border-white/15 rounded-2xl backdrop-blur-xl shadow-2xl hover:bg-white/15 transition-all duration-300 overflow-hidden flex flex-col"
          >
            <div class="relative h-48 overflow-hidden">
              <img
                :src="`https://img.youtube.com/vi/${teaching.id}/hqdefault.jpg`"
                :alt="teaching.title"
                class="w-full h-full object-cover transition-transform duration-500 hover:scale-110"
              />
              <div
                class="absolute inset-0 bg-gradient-to-t from-black/55 via-transparent to-transparent"
              ></div>
              <div
                class="absolute bottom-4 left-4 inline-flex items-center px-3 py-1 bg-white text-gray-900 rounded-full text-xs font-semibold tracking-wide shadow"
              >
                {{ teachings.badge }}
              </div>
            </div>
            <div class="p-6 flex flex-col flex-1">
              <h3
                class="text-lg font-semibold text-white mb-2 leading-snug line-clamp-2"
              >
                {{ teaching.title }}
              </h3>
              <p
                class="text-sm text-white/70 mb-5 leading-relaxed line-clamp-3 flex-1"
              >
                {{ teaching.description }}
              </p>
              <div
                class="flex items-center justify-between border-t border-white/20 pt-4"
              >
                <span class="text-xs uppercase tracking-[0.3em] text-white/50"
                  >YouTube</span
                >
                <a
                  :href="teaching.url"
                  target="_blank"
                  rel="noopener noreferrer"
                  class="inline-flex items-center text-sm font-semibold text-brand-orange hover:text-white transition-colors"
                >
                  {{ teachings.watchLabel }}
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
                </a>
              </div>
            </div>
          </div>
        </div>

        <div
          v-if="filteredVideos.length > videosPerPage"
          class="mt-8 flex items-center justify-center gap-3"
        >
          <button
            type="button"
            class="px-4 py-2 rounded-full border border-white/30 text-white/80 hover:text-white hover:border-white/60 disabled:opacity-40 disabled:cursor-not-allowed"
            :disabled="currentVideoPage === 1"
            @click="prevVideoPage"
          >
            Prev
          </button>
          <div class="flex items-center gap-2">
            <button
              v-for="page in totalVideoPages"
              :key="page"
              type="button"
              class="w-9 h-9 rounded-full border border-white/20 text-sm font-semibold"
              :class="
                page === currentVideoPage
                  ? 'bg-white text-brand-blue border-white'
                  : 'text-white/70 hover:text-white hover:border-white/50'
              "
              @click="goToVideoPage(page)"
            >
              {{ page }}
            </button>
          </div>
          <button
            type="button"
            class="px-4 py-2 rounded-full border border-white/30 text-white/80 hover:text-white hover:border-white/60 disabled:opacity-40 disabled:cursor-not-allowed"
            :disabled="currentVideoPage === totalVideoPages"
            @click="nextVideoPage"
          >
            Next
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, ref, watch, onMounted } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import heroImageCommunity from "@/assets/background/about-bg.jpg";
import heroImageTeaching from "@/assets/background/hero-bg.jpg";
import heroImageServe from "@/assets/background/bg-hero2.jpg";
import heroImageYouth from "@/assets/background/bg-hero3.jpg";
import { videoStore } from "@/services/videoStore";

export default {
  name: "Home",
  setup() {
    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const sundayBackground =
      "radial-gradient(circle at top left, rgba(255,164,91,0.18), transparent 50%), radial-gradient(circle at bottom right, rgba(34,112,130,0.15), transparent 45%)";
    const wednesdayBackground =
      "radial-gradient(circle at top right, rgba(34,112,130,0.2), transparent 50%), radial-gradient(circle at bottom left, rgba(255,164,91,0.15), transparent 40%)";

    const translations = {
      en: {
        heroSlides: [
          {
            id: "community",
            eyebrow: "Calvary Chapel Phnom Penh",
            title:
              "A worshiping community rooted in Scripture and present for Phnom Penh.",
            description:
              "We gather to teach verse by verse through the Bible, cultivate authentic relationships, and serve the city with the hope of Jesus.",
            image: heroImageCommunity,
            primaryCta: { label: "About Us", route: "/about" },
            secondaryCta: { label: "Contact Us", route: "/contact" },
          },
          {
            id: "teaching",
            eyebrow: "Verse-by-verse teaching",
            title: "Teaching the whole Bible with clarity and conviction.",
            description:
              "Join us Sunday mornings at 9:00 a.m. as we open the Scriptures, worship together, and practice the way of Jesus.",
            image: heroImageTeaching,
            primaryCta: { label: "Plan a Visit", route: "/contact" },
            secondaryCta: {
              label: "What We Believe",
              route: "/about/what-we-believe",
            },
          },
          {
            id: "serve",
            eyebrow: "Mercy & mission",
            title: "Serving Phnom Penh with meals, prayer, and practical care.",
            description:
              "From city outreaches to regional partnerships, teams are loving neighbors with tangible compassion in Jesus’ name.",
            image: heroImageServe,
            primaryCta: {
              label: "Explore Outreaches",
              route: "/ministries/outreaches",
            },
            secondaryCta: { label: "How to Help", route: "/how-to-help" },
          },
          {
            id: "prayer",
            eyebrow: "Citywide prayer",
            title: "Friday prayer night at 7:00 p.m.",
            description:
              "Gather with us every Friday at 7:00 p.m. to seek Jesus together in worship and intercession for Phnom Penh.",
            image: heroImageYouth,
            primaryCta: {
              label: "Join Prayer Night",
              route: "/events/upcoming-events",
            },
            secondaryCta: { label: "Contact Us", route: "/contact" },
          },
        ],
        weekly: {
          eyebrow: "Weekly Rhythm",
          title: "Regular Study & Worship",
          description:
            "Two anchor gatherings each week keep us rooted in Scripture and connected as a family.",
          gatherings: [
            {
              id: "sunday",
              accent: "blue",
              backgroundStyle: sundayBackground,
              dayLabel: "Sunday",
              time: "9:00",
              meridiem: "a.m.",
              badge: "Sunday Worship",
              title: "Sunday Morning Worship",
              body: "Verse-by-verse teaching, Spirit-led worship, and community for every generation gathered in the sanctuary each Sunday morning at 9:00 a.m.",
              ctaLabel: "Contact Us",
              ctaRoute: "/contact",
            },
            {
              id: "wednesday",
              accent: "orange",
              backgroundStyle: wednesdayBackground,
              dayLabel: "Wednesday",
              time: "7:00",
              meridiem: "p.m.",
              badge: "Wednesday Church",
              title: "Wednesday Night Church",
              body: "Wednesday night church family gathering featuring Scripture, discussion, and prayer beginning promptly at 7:00 p.m.",
              ctaLabel: "Contact Us",
              ctaRoute: "/events/calendar",
            },
          ],
        },
        upcoming: {
          eyebrow: "Upcoming",
          title: "What’s ahead this season",
          description:
            "A glimpse at the moments, gatherings, and opportunities we’re preparing for together.",
          ctaLabel: "More Details",
          items: [
            {
              id: "prayer-night",
              month: "Mar",
              day: "02",
              weekday: "Sat",
              title: "City Prayer Night",
              description:
                "An evening of prayer and worship for Phnom Penh, hosted in the main sanctuary at 6:30 PM.",
              route: "/events",
            },
            {
              id: "foundations-class",
              month: "Mar",
              day: "10",
              weekday: "Sun",
              title: "Foundations Class",
              description:
                "Four weeks exploring what we believe, how to grow, and where to serve. Sundays at 11:15 AM.",
              route: "/ministries",
            },
            {
              id: "serve-phnom-penh",
              month: "Mar",
              day: "16",
              weekday: "Sat",
              title: "Serve Phnom Penh",
              description:
                "Neighborhood outreach day providing meals, prayer, and practical care. Teams launch at 8:00 AM.",
              route: "/ministries/outreaches",
            },
          ],
        },
        teachings: {
          eyebrow: "Latest Teachings",
          title: "Recent messages from our pulpit",
          description:
            "Teachings captured each week to keep you connected to the verse-by-verse journey through Scripture.",
          badge: "Teaching Series",
          watchLabel: "Watch now",
          searchPlaceholder: "Search videos...",
          list: [
            {
              id: "IrTzMEBtbaI",
              title: "Latest Teaching 1",
              description: "Watch our latest message and be encouraged.",
              url: "https://www.youtube.com/watch?v=IrTzMEBtbaI",
            },
            {
              id: "AdDxynFrn8E",
              title: "Latest Teaching 2",
              description: "Join us for this week’s teaching.",
              url: "https://www.youtube.com/watch?v=AdDxynFrn8E",
            },
            {
              id: "FC0Tm4bNPkk",
              title: "Latest Teaching 3",
              description:
                "Stay connected with our community through God’s Word.",
              url: "https://www.youtube.com/watch?v=FC0Tm4bNPkk",
            },
          ],
        },
      },
      kh: {
        heroSlides: [
          {
            id: "community",
            eyebrow: "Calvary Chapel Phnom Penh",
            title:
              "សហគមន៍ថ្វាយបង្គំដែលជិតស្និទ្ធនឹងព្រះវចនៈ និងស្ថិតសម្រាប់ទីក្រុងភ្នំពេញ។",
            description:
              "យើងជួបជុំគ្នាដើម្បីបង្រៀនតាមជួរព្រះគម្ពីរ បង្កើតទំនាក់ទំនងពិតប្រាកដ និងបម្រើទីក្រុងដោយសេចក្តីសង្ឃឹមនៅក្នុងព្រះយេស៊ូវ។",
            image: heroImageCommunity,
            primaryCta: { label: "អំពីពួកយើង", route: "/about" },
            secondaryCta: { label: "ទាក់ទងមកកាន់ពួកយើង", route: "/contact" },
          },
          {
            id: "teaching",
            eyebrow: "ការបង្រៀនតាមជួរ",
            title: "បង្រៀនព្រះគម្ពីរទាំងស្រុងដោយច្បាស់លាស់ និងមុតមាំ។",
            description:
              "ចូលរួមជាមួយពួកយើងរៀងរាល់ថ្ងៃអាទិត្យ ម៉ោង 9 ព្រឹក ខណៈពេលយើងបើកព្រះគម្ពីរ ថ្វាយបង្គំ និងអនុវត្តផ្លូវព្រះយេស៊ូវ។",
            image: heroImageTeaching,
            primaryCta: { label: "រៀបចំការមកចូលរួម", route: "/contact" },
            secondaryCta: {
              label: "អ្វីដែលយើងជឿ",
              route: "/about/what-we-believe",
            },
          },
          {
            id: "serve",
            eyebrow: "មេត្តាករុណា និងភារកិច្ច",
            title:
              "បម្រើទីក្រុងភ្នំពេញដោយអាហារ ការអធិស្ឋាន និងការថែទាំជាក់ស្តែង។",
            description:
              "ចាប់ពីការចេញផ្សាយក្នុងទីក្រុងដល់ដៃគូតំបន់ ក្រុមកំពុងស្រឡាញ់អ្នកជិតខាងដោយមេត្តាករុណាជាក់ស្តែងក្នុងនាមព្រះយេស៊ូវ។",
            image: heroImageServe,
            primaryCta: {
              label: "ស្វែងរកការចេញផ្សាយ",
              route: "/ministries/outreaches",
            },
            secondaryCta: { label: "របៀបជួយ", route: "/how-to-help" },
          },
          {
            id: "prayer",
            eyebrow: "ការអធិស្ឋានទាំងទីក្រុង",
            title: "យប់អធិស្ឋានថ្ងៃសុក្រ ម៉ោង 7 យប់។",
            description:
              "ជួបជាមួយយើងរៀងរាល់ថ្ងៃសុក្រ ម៉ោង 7 យប់ ដើម្បីស្វែងរកព្រះយេស៊ូវរួមគ្នា ក្នុងការថ្វាយបង្គំ និងការអធិស្ឋានសម្រាប់ទីក្រុងភ្នំពេញ។",
            image: heroImageYouth,
            primaryCta: {
              label: "ចូលរួមយប់អធិស្ឋាន",
              route: "/events/upcoming-events",
            },
            secondaryCta: { label: "ទាក់ទងមកកាន់ពួកយើង", route: "/contact" },
          },
        ],
        weekly: {
          eyebrow: "កាលវិភាគប្រចាំសប្ដាហ៍",
          title: "ការសិក្សា និងការថ្វាយបង្គំទៀងទាត់",
          description:
            "កិច្ចប្រជុំសំខាន់ពីររៀងរាល់សប្ដាហ៍ជួយឲ្យយើងជាប់នឹងព្រះវចនៈ និងកាន់តែជាប់ស្និទ្ធជាគ្រួសារ។",
          gatherings: [
            {
              id: "sunday",
              accent: "blue",
              backgroundStyle: sundayBackground,
              dayLabel: "ថ្ងៃអាទិត្យ",
              time: "9:00",
              meridiem: "ព្រឹក",
              badge: "ការថ្វាយបង្គំថ្ងៃអាទិត្យ",
              title: "ការថ្វាយបង្គំនៅព្រឹកថ្ងៃអាទិត្យ",
              body: "បង្រៀនតាមជួរ ថ្វាយបង្គំដោយព្រះវិញ្ញាណ និងសហគមន៍សម្រាប់គ្រួសារទាំងអស់ នៅសាលធំរៀងរាល់ព្រឹកថ្ងៃអាទិត្យ ម៉ោង 9:00។",
              ctaLabel: "ទាក់ទងមកកាន់ពួកយើង",
              ctaRoute: "/contact",
            },
            {
              id: "wednesday",
              accent: "orange",
              backgroundStyle: wednesdayBackground,
              dayLabel: "ថ្ងៃពុធ",
              time: "7:00",
              meridiem: "យប់",
              badge: "សិក្សាព្រះគម្ពីរ ថ្ងៃពុធ",
              title: "សិក្សាព្រះគម្ពីរ យប់ថ្ងៃពុធ",
              body: "ការជួបជុំគ្រួសារយប់ថ្ងៃពុធ ដែលមានព្រះគម្ពីរ ការពិភាក្សា និងការអធិស្ឋាន ចាប់ផ្តើមម៉ោង 7:00 ពេលល្ងាច។",
              ctaLabel: "ទាក់ទងមកកាន់ពួកយើង",
              ctaRoute: "/events/calendar",
            },
          ],
        },
        upcoming: {
          eyebrow: "ព្រឹត្តិការណ៍ជិតមកដល់",
          title: "អ្វីខ្លះនៅមុខរដូវនេះ",
          description:
            "រូបភាពខ្លីៗនៃខណៈពេល ការជួបជុំ និងឱកាសដែលយើងកំពុងរៀបចំរួមគ្នា។",
          ctaLabel: "ព័ត៌មានបន្ថែម",
          items: [
            {
              id: "prayer-night",
              month: "មីនា",
              day: "02",
              weekday: "សៅរ៍",
              title: "យប់អធិស្ឋានទីក្រុង",
              description:
                "យប់អធិស្ឋាន និងការថ្វាយបង្គំសម្រាប់ទីក្រុងភ្នំពេញ នៅសាលធំ ម៉ោង 6:30 ល្ងាច។",
              route: "/events",
            },
            {
              id: "foundations-class",
              month: "មីនា",
              day: "10",
              weekday: "អាទិត្យ",
              title: "ថ្នាក់មូលដ្ឋាន",
              description:
                "៤ សប្ដាហ៍សិក្សាអំពីអ្វីដែលយើងជឿ របៀបលូតលាស់ និងទីកន្លែងបម្រើ រៀងរាល់ថ្ងៃអាទិត្យ ម៉ោង 11:15 ព្រឹក។",
              route: "/ministries",
            },
            {
              id: "serve-phnom-penh",
              month: "មីនា",
              day: "16",
              weekday: "សៅរ៍",
              title: "បម្រើទីក្រុងភ្នំពេញ",
              description:
                "ថ្ងៃចេញផ្សាយក្នុងសង្កាត់ ផ្តល់អាហារ ការអធិស្ឋាន និងការថែទាំជាក់ស្តែង ក្រុមចេញដំណើរម៉ោង 8 ព្រឹក។",
              route: "/ministries/outreaches",
            },
          ],
        },
        teachings: {
          eyebrow: "ការបង្រៀនថ្មីៗ",
          title: "សារថ្មីៗពីឆាកសន្មត",
          description:
            "ការបង្រៀនត្រូវបានថតរៀងរាល់សប្ដាហ៍ ដើម្បីអោយអ្នកនៅតែភ្ជាប់ជាមួយដំណើរតាមជួរព្រះគម្ពីរ។",
          badge: "ស៊េរីការបង្រៀន",
          watchLabel: "មើលឥឡូវនេះ",
          searchPlaceholder: "ស្វែងរកវីដេអូ...",
          list: [
            {
              id: "IrTzMEBtbaI",
              title: "ការបង្រៀនថ្មី ទី១",
              description: "មើលសារថ្មីបំផុតរបស់យើង ហើយទទួលបានការលើកទឹកចិត្ត។",
              url: "https://www.youtube.com/watch?v=IrTzMEBtbaI",
            },
            {
              id: "AdDxynFrn8E",
              title: "ការបង្រៀនថ្មី ទី២",
              description: "ចូលរួមស្តាប់ការបង្រៀនសប្ដាហ៍នេះជាមួយយើង។",
              url: "https://www.youtube.com/watch?v=AdDxynFrn8E",
            },
            {
              id: "FC0Tm4bNPkk",
              title: "ការបង្រៀនថ្មី ទី៣",
              description: "នៅតែភ្ជាប់ជាមួយសហគមន៍តាមរយៈព្រះបន្ទូលរបស់ព្រះ។",
              url: "https://www.youtube.com/watch?v=FC0Tm4bNPkk",
            },
          ],
        },
      },
    };

    const currentSlideIndex = ref(0);

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    const heroSlides = computed(() => content.value.heroSlides);
    const activeSlide = computed(
      () => heroSlides.value[currentSlideIndex.value],
    );

    const heroBackgroundStyle = computed(() => ({
      backgroundImage: `linear-gradient(120deg, rgba(17, 39, 84, 0.25), rgba(17, 39, 84, 0.45)), url(${activeSlide.value.image})`,
    }));

    const weekly = computed(() => content.value.weekly);
    const weeklyGatherings = computed(() => weekly.value.gatherings);
    const upcoming = computed(() => content.value.upcoming);
    const upcomingItems = computed(() => upcoming.value.items);
    const teachings = computed(() => content.value.teachings);

    const storedVideos = ref([]);
    const videosPerPage = 3;
    const currentVideoPage = ref(1);
    const videoSearchQuery = ref("");

    const allVideos = computed(() => {
      if (storedVideos.value.length) {
        return storedVideos.value.map((video) => ({
          id: video.youtubeId,
          title: video.title,
          description: video.description,
          url:
            video.youtubeUrl ||
            `https://www.youtube.com/watch?v=${video.youtubeId}`,
        }));
      }
      return teachings.value.list;
    });

    const filteredVideos = computed(() => {
      if (!videoSearchQuery.value.trim()) {
        return allVideos.value;
      }
      const query = videoSearchQuery.value.toLowerCase().trim();
      return allVideos.value.filter(
        (video) =>
          video.title.toLowerCase().includes(query) ||
          video.description.toLowerCase().includes(query),
      );
    });

    const totalVideoPages = computed(() =>
      filteredVideos.value.length
        ? Math.ceil(filteredVideos.value.length / videosPerPage)
        : 1,
    );

    const displayedTeachings = computed(() => {
      const start = (currentVideoPage.value - 1) * videosPerPage;
      return filteredVideos.value.slice(start, start + videosPerPage);
    });

    const performVideoSearch = () => {
      currentVideoPage.value = 1;
    };

    const clearVideoSearch = () => {
      videoSearchQuery.value = "";
      currentVideoPage.value = 1;
    };

    const nextSlide = () => {
      currentSlideIndex.value =
        (currentSlideIndex.value + 1) % heroSlides.value.length;
    };

    const prevSlide = () => {
      currentSlideIndex.value =
        (currentSlideIndex.value - 1 + heroSlides.value.length) %
        heroSlides.value.length;
    };

    const goToSlide = (index) => {
      currentSlideIndex.value = index;
    };

    watch(selectedLanguage, () => {
      currentSlideIndex.value = 0;
    });

    onMounted(() => {
      storedVideos.value = videoStore.getAll();
      if (currentVideoPage.value > totalVideoPages.value) {
        currentVideoPage.value = 1;
      }
    });

    watch(
      () => storedVideos.value.length,
      () => {
        if (currentVideoPage.value > totalVideoPages.value) {
          currentVideoPage.value = 1;
        }
      },
    );

    watch(videoSearchQuery, () => {
      if (currentVideoPage.value > totalVideoPages.value) {
        currentVideoPage.value = 1;
      }
    });

    const nextVideoPage = () => {
      if (currentVideoPage.value < totalVideoPages.value) {
        currentVideoPage.value += 1;
      }
    };

    const prevVideoPage = () => {
      if (currentVideoPage.value > 1) {
        currentVideoPage.value -= 1;
      }
    };

    const goToVideoPage = (page) => {
      currentVideoPage.value = page;
    };

    return {
      heroSlides,
      activeSlide,
      heroBackgroundStyle,
      nextSlide,
      prevSlide,
      goToSlide,
      currentSlideIndex,
      weekly,
      weeklyGatherings,
      upcoming,
      upcomingItems,
      teachings,
      storedVideos,
      displayedTeachings,
      videosPerPage,
      currentVideoPage,
      totalVideoPages,
      nextVideoPage,
      prevVideoPage,
      goToVideoPage,
      videoSearchQuery,
      filteredVideos,
      performVideoSearch,
      clearVideoSearch,
    };
  },
};
</script>
