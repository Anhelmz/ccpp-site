<template>
  <div class="relative isolate overflow-hidden">
    <div class="pointer-events-none absolute inset-0 -z-10">
      <div
        class="absolute left-1/2 top-[-8rem] h-96 w-96 -translate-x-1/2 rounded-full bg-primary-400/40 blur-3xl sm:h-[32rem] sm:w-[32rem]"
      ></div>
      <div
        class="absolute bottom-[-10rem] right-[-6rem] h-[28rem] w-[28rem] rounded-full bg-primary-700/30 blur-3xl"
      ></div>
    </div>

    <section
      class="relative flex min-h-screen flex-col px-4 py-20 sm:px-8 lg:px-12"
    >
      <div
        class="pointer-events-none absolute inset-0 -z-10 bg-gradient-to-br from-slate-950 via-slate-900 to-primary-950"
      ></div>
      <div class="pointer-events-none absolute inset-0 -z-10 opacity-60">
        <div
          class="absolute left-[-6rem] top-[-6rem] h-80 w-80 rounded-full bg-primary-500/40 blur-3xl"
        ></div>
        <div
          class="absolute bottom-[-8rem] right-[-4rem] h-96 w-96 rounded-full bg-primary-400/20 blur-3xl"
        ></div>
      </div>

      <div class="w-full text-slate-100">
        <slot />
      </div>
    </section>
  </div>
</template>
