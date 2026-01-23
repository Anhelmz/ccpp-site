<template>
  <div class="bg-white">
    <!-- Hero -->
    <section
      class="relative min-h-[55vh] flex items-center overflow-hidden pt-24 md:pt-28 pb-16 bg-gradient-to-br from-[#0f172a] via-[#11263f] to-[#1d3a54] text-white"
    >
      <div class="absolute inset-0" aria-hidden="true">
        <div
          class="absolute inset-0 opacity-30 mix-blend-screen"
          style="
            background-image:
              radial-gradient(
                circle at 20% 30%,
                rgba(255, 255, 255, 0.25),
                transparent 45%
              ),
              radial-gradient(
                circle at 70% 15%,
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
          class="absolute inset-y-[-35%] left-[-15%] w-1/2 bg-gradient-to-r from-brand-orange/30 via-transparent to-transparent blur-[150px] opacity-70"
        ></div>
        <div
          class="absolute inset-y-[-35%] right-[-15%] w-2/3 bg-gradient-to-l from-brand-blue/30 via-transparent to-transparent blur-[200px] opacity-70"
        ></div>
        <div
          class="absolute inset-x-0 bottom-0 h-40 bg-gradient-to-t from-black/70 via-black/40 to-transparent"
        ></div>
      </div>
      <div
        class="relative z-10 w-full max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center space-y-6"
      >
        <p class="text-sm uppercase tracking-[0.35em] text-white/70">
          {{ hero.eyebrow }}
        </p>
        <h1 class="text-4xl md:text-5xl lg:text-6xl font-bold leading-tight">
          {{ hero.title }}
        </h1>
        <p
          class="text-lg md:text-xl text-white/85 max-w-3xl mx-auto font-light"
        >
          {{ hero.subtitle }}
        </p>
        <div class="flex justify-center gap-4 pt-4 text-white/80">
          <a
            href="https://www.facebook.com/calvarychapelpp"
            target="_blank"
            rel="noopener noreferrer"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="Facebook"
          >
            <font-awesome-icon :icon="['fab', 'facebook']" class="w-5 h-5" />
          </a>
          <a
            href="mailto:calvarychapelphnompenh@gmail.com"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="Email"
          >
            <font-awesome-icon :icon="['fas', 'envelope']" class="w-5 h-5" />
          </a>
          <a
            href="https://www.youtube.com/@CalvaryChapelPhnomPenh"
            target="_blank"
            rel="noopener noreferrer"
            class="w-11 h-11 rounded-full border border-white/30 flex items-center justify-center hover:bg-white/10 transition"
            aria-label="YouTube"
          >
            <font-awesome-icon :icon="['fab', 'youtube']" class="w-5 h-5" />
          </a>
        </div>
      </div>
    </section>

    <!-- Info + Form -->
    <section
      class="py-16 md:py-20 bg-gradient-to-b from-white via-gray-50 to-gray-100"
    >
      <div
        class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid gap-10 lg:grid-cols-2"
      >
        <div class="space-y-8">
          <div
            class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-6"
          >
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">
              {{ info.eyebrow }}
            </p>
            <h2 class="text-3xl font-semibold text-gray-900">
              {{ info.title }}
            </h2>
            <p class="text-gray-600">
              {{ info.description }}
            </p>
            <div class="space-y-6">
              <div class="flex gap-4">
                <div>
                  <h3 class="font-semibold text-gray-900">
                    {{ info.address.label }}
                  </h3>
                  <p class="text-gray-600 leading-relaxed">
                    <span v-for="(line, idx) in info.address.lines" :key="idx">
                      {{ line }}<br />
                    </span>
                  </p>
                </div>
              </div>
              <div class="flex gap-4">
                <div>
                  <h3 class="font-semibold text-gray-900">
                    {{ info.phone.label }}
                  </h3>
                  <p class="text-gray-600">{{ info.phone.value }}</p>
                </div>
              </div>
              <div class="flex gap-4">
                <div>
                  <h3 class="font-semibold text-gray-900">
                    {{ info.email.label }}
                  </h3>
                  <p class="text-gray-600">{{ info.email.value }}</p>
                </div>
              </div>
            </div>
          </div>

          <div
            class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8 space-y-4"
          >
            <p class="text-xs uppercase tracking-[0.35em] text-brand-orange/80">
              {{ gatherings.eyebrow }}
            </p>
            <h3 class="text-2xl font-semibold text-gray-900">
              {{ gatherings.title }}
            </h3>
            <ul class="space-y-4 text-gray-600">
              <li
                v-for="(item, idx) in gatherings.items"
                :key="item.id"
                class="flex items-center justify-between"
                :class="idx === 0 ? 'border-b border-gray-100 pb-3' : ''"
              >
                <div>
                  <p class="font-semibold text-gray-900">{{ item.title }}</p>
                  <p class="text-sm text-gray-500">{{ item.location }}</p>
                </div>
                <span class="font-medium text-gray-700">{{ item.time }}</span>
              </li>
            </ul>
          </div>
        </div>

        <div
          class="rounded-[28px] border border-gray-100 bg-white shadow-2xl p-8"
        >
          <div class="space-y-2 mb-6 text-center lg:text-left">
            <p class="text-xs uppercase tracking-[0.35em] text-brand-blue/70">
              {{ formTexts.eyebrow }}
            </p>
            <h2 class="text-3xl font-semibold text-gray-900">
              {{ formTexts.title }}
            </h2>
            <p class="text-gray-600">{{ formTexts.subtitle }}</p>
          </div>
          <form class="space-y-6" @submit.prevent="submitForm">
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">{{
                  formTexts.fields.name
                }}</label>
                <input
                  v-model="form.name"
                  type="text"
                  required
                  class="input-field"
                  :placeholder="formTexts.placeholders.name"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">{{
                  formTexts.fields.email
                }}</label>
                <input
                  v-model="form.email"
                  type="email"
                  required
                  class="input-field"
                  :placeholder="formTexts.placeholders.email"
                />
              </div>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">{{
                  formTexts.fields.phone
                }}</label>
                <input
                  v-model="form.phone"
                  type="tel"
                  class="input-field"
                  :placeholder="formTexts.placeholders.phone"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">{{
                  formTexts.fields.reason
                }}</label>
                <select v-model="form.reason" required class="input-field">
                  <option value="">{{ formTexts.placeholders.reason }}</option>
                  <option
                    v-for="option in formTexts.reasonOptions"
                    :key="option.value"
                    :value="option.value"
                  >
                    {{ option.label }}
                  </option>
                </select>
              </div>
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">{{
                formTexts.fields.message
              }}</label>
              <textarea
                v-model="form.message"
                required
                rows="5"
                class="input-field"
                :placeholder="formTexts.placeholders.message"
              ></textarea>
            </div>
            <button
              type="submit"
              :disabled="submitting"
              class="btn-primary w-full"
              :class="{ 'opacity-50 cursor-not-allowed': submitting }"
            >
              {{ submitting ? formTexts.sending : formTexts.submit }}
            </button>
          </form>
          <div
            v-if="message"
            class="mt-6 rounded-2xl p-4 border"
            :class="
              messageType === 'success'
                ? 'bg-green-50 text-green-700 border-green-100'
                : 'bg-red-50 text-red-700 border-red-100'
            "
          >
            {{ message }}
          </div>
        </div>
      </div>
    </section>

    <!-- Map -->
    <section class="py-16 bg-white">
      <div
        class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 grid gap-10 lg:grid-cols-2 items-center"
      >
        <div class="space-y-4">
          <p class="text-sm uppercase tracking-[0.35em] text-gray-400">
            {{ map.eyebrow }}
          </p>
          <h2 class="text-3xl font-semibold text-gray-900">{{ map.title }}</h2>
          <p class="text-gray-600">
            {{ map.description }}
          </p>
          <a
            href="https://maps.app.goo.gl/a4YGCcvWRLy6BvRz9"
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex items-center text-brand-blue font-semibold hover:text-brand-orange transition-colors"
          >
            {{ map.cta }}
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
        <div
          class="rounded-[32px] overflow-hidden border border-gray-100 shadow-2xl"
        >
          <iframe
            src="https://www.google.com/maps?q=Street+26BT,+House+428,+Phnom+Penh,+Cambodia&output=embed"
            width="100%"
            height="420"
            style="border: 0"
            allowfullscreen=""
            loading="lazy"
            referrerpolicy="no-referrer-when-downgrade"
            class="w-full"
          ></iframe>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { ref, computed } from "vue";
import { storeToRefs } from "pinia";
import { useLanguageStore } from "@/stores/language";
import { contactRequestService } from "@/services/contactRequestService";

export default {
  name: "Contact",
  setup() {
    const form = ref({
      name: "",
      email: "",
      phone: "",
      reason: "",
      message: "",
    });

    const submitting = ref(false);
    const message = ref("");
    const messageType = ref("");

    const { selectedLanguage } = storeToRefs(useLanguageStore());

    const translations = {
      en: {
        hero: {
          eyebrow: "We'd love to hear from you",
          title: "Contact Calvary Chapel Phnom Penh",
          subtitle:
            "Questions, prayer requests, or planning a visit? Reach out and our team will respond quickly.",
        },
        info: {
          eyebrow: "Visit & connect",
          title: "Get in touch",
          description:
            "Our team is available throughout the week to answer questions, pray with you, and help you plug in.",
          address: {
            label: "Address",
            lines: [
              "Street 26BT, House 428",
              "Boeung Tompun",
              "Phnom Penh, Cambodia",
            ],
          },
          phone: { label: "Phone", value: "+855 15 81 44 40" },
          email: { label: "Email", value: "calvarychapelphnompenh@gmail.com" },
        },
        gatherings: {
          eyebrow: "Gatherings",
          title: "Regular Service",
          items: [
            {
              id: "sun",
              title: "Sunday Bible Study",
              location: "Main sanctuary",
              time: "9:00 AM",
            },
            {
              id: "wed",
              title: "Wednesday Bible Study",
              location: "Mid-week gathering",
              time: "7:00 PM",
            },
          ],
        },
        form: {
          eyebrow: "Reach out",
          title: "Send us a message",
          subtitle: "We typically respond within one business day.",
          fields: {
            name: "Full Name *",
            email: "Email *",
            phone: "Phone",
            reason: "Why are you visiting? *",
            message: "Message *",
          },
          placeholders: {
            name: "Your full name",
            email: "your.email@example.com",
            phone: "+855 15 81 44 40",
            reason: "Select one",
            message: "How can we help you?",
          },
          reasonOptions: [
            { value: "Plan a visit", label: "Plan a visit" },
            { value: "Prayer request", label: "Prayer request" },
            { value: "Counseling / care", label: "Counseling / care" },
            { value: "Serving / volunteer", label: "Serving / volunteer" },
            {
              value: "Membership / next steps",
              label: "Membership / next steps",
            },
            { value: "General question", label: "General question" },
            { value: "Other", label: "Other" },
          ],
          submit: "Send Message",
          sending: "Sending...",
          success: "Thank you for your message! We'll get back to you soon.",
          error:
            "Sorry, there was an error sending your message. Please try again.",
        },
        map: {
          eyebrow: "Visit us",
          title: "Sunday gatherings in Phnom Penh",
          description:
            "We meet just south of central Phnom Penh near Boeung Tompun. Use the map for directions or open the route in Google Maps.",
          cta: "Open directions",
        },
        cta: { primary: "Contact Us" },
        footerNote:
          "For detailed schedules and new events, check the calendar or contact us—new opportunities are added regularly.",
      },
      kh: {
        hero: {
          eyebrow: "យើងចង់ស្តាប់ពីអ្នក",
          title: "ទាក់ទងព្រះវិហារកាល់វ៉ារីឆាបផលភ្នំពេញ",
          subtitle:
            "សំណួរ សំណូមពរអធិស្ឋាន ឬរៀបចំការមកទស្សនា? សូមទាក់ទង មន្ត្រីរបស់យើងនឹងឆ្លើយតបរហ័ស។",
        },
        info: {
          eyebrow: "មកទស្សនា និងតភ្ជាប់",
          title: "ទាក់ទងមកកាន់ពួកយើង",
          description:
            "ក្រុមរបស់យើងមានស្រាប់ទូទាំងសប្ដាហ៍ ដើម្បីឆ្លើយសំណួរ អធិស្ឋានជាមួយអ្នក និងជួយអ្នករួមចំណែក។",
          address: {
            label: "អាសយដ្ឋាន",
            lines: [
              "ផ្លូវ 26BT ផ្ទះលេខ 428",
              "បឹងទំពុន",
              "ភ្នំពេញ ប្រទេសកម្ពុជា",
            ],
          },
          phone: { label: "ទូរស័ព្ទ", value: "+855 15 81 44 40" },
          email: { label: "អ៊ីមែល", value: "calvarychapelphnompenh@gmail.com" },
        },
        gatherings: {
          eyebrow: "កាលវិភាគសេវាកម្ម",
          title: "សេវាប្រចាំ",
          items: [
            {
              id: "sun",
              title: "សិក្សាព្រះគម្ពីរ ថ្ងៃអាទិត្យ",
              location: "សាលធំ",
              time: "9:00 ព្រឹក",
            },
            {
              id: "wed",
              title: "សិក្សាព្រះគម្ពីរ ថ្ងៃពុធ",
              location: "ការជួបកណ្ដាលសប្ដាហ៍",
              time: "7:00 ល្ងាច",
            },
          ],
        },
        form: {
          eyebrow: "ទាក់ទង",
          title: "ផ្ញើសារមកយើង",
          subtitle: "យើងភាគច្រើនឆ្លើយតបក្នុងរយៈពេលមួយថ្ងៃធ្វើការ។",
          fields: {
            name: "ឈ្មោះពេញ *",
            email: "អ៊ីមែល *",
            phone: "ទូរស័ព្ទ",
            reason: "ហេតុអ្វីបានមកទស្សនា? *",
            message: "សារ *",
          },
          placeholders: {
            name: "ឈ្មោះពេញរបស់អ្នក",
            email: "your.email@example.com",
            phone: "+855 15 81 44 40",
            reason: "ជ្រើសរើសមួយ",
            message: "យើងអាចជួយអ្វីដល់អ្នក?",
          },
          reasonOptions: [
            { value: "Plan a visit", label: "រៀបចំការមកទស្សនា" },
            { value: "Prayer request", label: "សំណូមពរអធិស្ឋាន" },
            { value: "Counseling / care", label: "ការប្រឹក្សា / ថែទាំ" },
            { value: "Serving / volunteer", label: "បម្រើ / ស្ម័គ្រចិត្ត" },
            {
              value: "Membership / next steps",
              label: "សមាជិកភាព / ជំហានបន្ទាប់",
            },
            { value: "General question", label: "សំណួរទូទៅ" },
            { value: "Other", label: "ផ្សេងៗ" },
          ],
          submit: "ផ្ញើសារ",
          sending: "កំពុងផ្ញើ...",
          success: "សូមអរគុណសម្រាប់សាររបស់អ្នក! យើងនឹងតបត្រូវឆាប់ៗនេះ។",
          error: "សូមទោស មានបញ្ហា​ក្នុងការផ្ញើសារ សូមព្យាយាមម្តងទៀត។",
        },
        map: {
          eyebrow: "មកទស្សនាយើង",
          title: "ការជួបជុំថ្ងៃអាទិត្យនៅភ្នំពេញ",
          description:
            "យើងជួបនៅខាងត្បូងមួយចំនួននៃទីក្រុងភ្នំពេញ នៅជិតបឹងទំពុន។ ប្រើផែនទីសម្រាប់ទិសដៅ ឬបើកផ្លូវក្នុង Google Maps។",
          cta: "បើកទិសដៅ",
        },
        cta: { primary: "ទាក់ទងមកកាន់ពួកយើង" },
        footerNote:
          "សម្រាប់កាលវិភាគលម្អិត និងព្រឹត្តិការណ៍ថ្មី សូមពិនិត្យប្រតិទិន ឬទាក់ទងមកយើង—ឱកាសថ្មីត្រូវបានបន្ថែមជាប្រចាំ។",
      },
    };

    const content = computed(
      () => translations[selectedLanguage.value] || translations.en,
    );
    const hero = computed(() => content.value.hero);
    const info = computed(() => content.value.info);
    const gatherings = computed(() => content.value.gatherings);
    const formTexts = computed(() => content.value.form);
    const map = computed(() => content.value.map);
    const cta = computed(() => content.value.cta);
    const footerNote = computed(() => content.value.footerNote);

    const submitForm = async () => {
      submitting.value = true;
      message.value = "";

      try {
        await contactRequestService.createContactRequest({
          name: form.value.name,
          email: form.value.email,
          phone: form.value.phone || undefined,
          reason: form.value.reason,
          message: form.value.message,
        });

        message.value = formTexts.value.success;
        messageType.value = "success";

        form.value = {
          name: "",
          email: "",
          phone: "",
          reason: "",
          message: "",
        };
      } catch (error) {
        message.value = error.message || formTexts.value.error;
        messageType.value = "error";
      } finally {
        submitting.value = false;
      }
    };

    const scheduleVisit = () => {
      alert(info.value.phone.value);
    };

    return {
      form,
      submitting,
      message,
      messageType,
      submitForm,
      scheduleVisit,
      hero,
      info,
      gatherings,
      formTexts,
      map,
      cta,
      footerNote,
    };
  },
};
</script>
