<template>
  <div class="px-4 sm:px-0">
    <!-- Hero Section -->
    <div class="bg-main text-white py-16 mb-12">
      <div class="max-w-4xl mx-auto text-center">
        <h1 class="text-4xl md:text-6xl font-bold mb-4">Support Our Mission</h1>
        <p class="text-xl md:text-2xl">Your Generosity Makes a Difference</p>
      </div>
    </div>

    <!-- Donation Options -->
    <div class="max-w-6xl mx-auto mb-12">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold text-gray-900 mb-4">Ways to Give</h2>
        <p class="text-lg text-gray-600 max-w-3xl mx-auto">
          Your support helps us continue our mission of spreading God's love
          through community service, missions, and spiritual growth programs.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8 mb-12">
        <!-- General Fund -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">⛪</div>
          <h3 class="text-xl font-semibold mb-3">General Fund</h3>
          <p class="text-gray-600 mb-6">
            Support our day-to-day operations, facilities, and general ministry
            activities.
          </p>
          <button class="btn-primary" @click="selectDonationType('general')">
            Donate Now
          </button>
        </div>

        <!-- Missions -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">🌍</div>
          <h3 class="text-xl font-semibold mb-3">Missions</h3>
          <p class="text-gray-600 mb-6">
            Help fund our international mission trips and global outreach
            programs.
          </p>
          <button class="btn-primary" @click="selectDonationType('missions')">
            Donate Now
          </button>
        </div>

        <!-- Youth Ministry -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">👨‍👩‍👧‍👦</div>
          <h3 class="text-xl font-semibold mb-3">Youth Ministry</h3>
          <p class="text-gray-600 mb-6">
            Support programs that nurture the spiritual growth of our young
            people.
          </p>
          <button class="btn-primary" @click="selectDonationType('youth')">
            Donate Now
          </button>
        </div>

        <!-- Community Outreach -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">🤝</div>
          <h3 class="text-xl font-semibold mb-3">Community Outreach</h3>
          <p class="text-gray-600 mb-6">
            Help us serve our local community through food banks, shelters, and
            support programs.
          </p>
          <button class="btn-primary" @click="selectDonationType('outreach')">
            Donate Now
          </button>
        </div>

        <!-- Building Fund -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">🏗️</div>
          <h3 class="text-xl font-semibold mb-3">Building Fund</h3>
          <p class="text-gray-600 mb-6">
            Contribute to facility improvements and expansion projects.
          </p>
          <button class="btn-primary" @click="selectDonationType('building')">
            Donate Now
          </button>
        </div>

        <!-- Memorial Fund -->
        <div class="card text-center hover:shadow-lg transition-shadow">
          <div class="text-4xl mb-4">💝</div>
          <h3 class="text-xl font-semibold mb-3">Memorial Fund</h3>
          <p class="text-gray-600 mb-6">
            Honor the memory of loved ones through special memorial
            contributions.
          </p>
          <button class="btn-primary" @click="selectDonationType('memorial')">
            Donate Now
          </button>
        </div>
      </div>
    </div>

    <!-- Donation Form Modal -->
    <div
      v-if="showDonationForm"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50"
    >
      <div
        class="relative top-20 mx-auto p-5 border w-96 max-w-md shadow-lg rounded-md bg-white"
      >
        <div class="mt-3">
          <div class="flex justify-between items-center mb-4">
            <h3 class="text-lg font-medium text-gray-900">
              Donate to {{ donationType }}
            </h3>
            <button
              class="text-gray-400 hover:text-gray-600"
              @click="closeDonationForm"
            >
              <span class="text-2xl">&times;</span>
            </button>
          </div>

          <form @submit.prevent="submitDonation">
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Amount
              </label>
              <div class="grid grid-cols-3 gap-2 mb-2">
                <button
                  v-for="amount in quickAmounts"
                  :key="amount"
                  type="button"
                  class="btn-secondary text-sm"
                  :class="{
                    'bg-main text-white': donationForm.amount === amount,
                  }"
                  @click="donationForm.amount = amount"
                >
                  ${{ amount }}
                </button>
              </div>
              <input
                v-model="donationForm.amount"
                type="number"
                min="1"
                step="0.01"
                class="input-field"
                placeholder="Enter amount"
              />
            </div>

            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Your Name
              </label>
              <input
                v-model="donationForm.name"
                type="text"
                class="input-field"
                placeholder="Your full name"
              />
            </div>

            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Email
              </label>
              <input
                v-model="donationForm.email"
                type="email"
                class="input-field"
                placeholder="your.email@example.com"
              />
            </div>

            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Message (Optional)
              </label>
              <textarea
                v-model="donationForm.message"
                rows="3"
                class="input-field"
                placeholder="Add a personal message..."
              ></textarea>
            </div>

            <div class="mb-6">
              <label class="flex items-center">
                <input
                  v-model="donationForm.isAnonymous"
                  type="checkbox"
                  class="rounded border-gray-300 text-main shadow-sm focus:border-secondary focus:ring focus:ring-secondary focus:ring-opacity-50"
                />
                <span class="ml-2 text-sm text-gray-600"
                  >Make this donation anonymous</span
                >
              </label>
            </div>

            <div class="flex justify-end space-x-3">
              <button
                type="button"
                class="btn-secondary"
                @click="closeDonationForm"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="submitting"
                class="btn-primary"
                :class="{ 'opacity-50 cursor-not-allowed': submitting }"
              >
                {{
                  submitting
                    ? "Processing..."
                    : `Donate $${donationForm.amount || "0"}`
                }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>

    <!-- Other Giving Methods -->
    <div class="bg-primary py-12 mb-12">
      <div class="max-w-4xl mx-auto">
        <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">
          Other Ways to Give
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div class="card">
            <h3 class="text-xl font-semibold mb-4">Online Giving</h3>
            <p class="text-gray-600 mb-4">
              Set up recurring donations or make one-time gifts through our
              secure online platform.
            </p>
            <ul class="text-gray-600 space-y-2">
              <li>• Credit/Debit Cards</li>
              <li>• Bank Transfer (ACH)</li>
              <li>• PayPal</li>
              <li>• Apple Pay / Google Pay</li>
            </ul>
          </div>

          <div class="card">
            <h3 class="text-xl font-semibold mb-4">Traditional Methods</h3>
            <p class="text-gray-600 mb-4">
              Prefer traditional giving methods? We accept these forms of
              donation as well.
            </p>
            <ul class="text-gray-600 space-y-2">
              <li>• Cash and Checks</li>
              <li>• Stock Donations</li>
              <li>• Estate Planning</li>
              <li>• In-Kind Donations</li>
            </ul>
          </div>
        </div>
      </div>
    </div>

    <!-- Impact Stories -->
    <div class="max-w-4xl mx-auto mb-12">
      <h2 class="text-3xl font-bold text-center text-gray-900 mb-8">
        Your Impact
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div class="card">
          <h3 class="text-xl font-semibold mb-3">Local Community Impact</h3>
          <p class="text-gray-600 mb-4">
            "Thanks to the generous donations from our congregation, we were
            able to provide meals for 500 families during the holiday season and
            support 50 children with school supplies."
          </p>
          <p class="text-sm text-gray-500">- Community Outreach Team</p>
        </div>

        <div class="card">
          <h3 class="text-xl font-semibold mb-3">Global Mission Impact</h3>
          <p class="text-gray-600 mb-4">
            "Your support helped us complete the water well project in Kenya,
            providing clean water to 300 families. The community is forever
            grateful for this life-changing gift."
          </p>
          <p class="text-sm text-gray-500">- Mission Team Kenya</p>
        </div>
      </div>
    </div>

    <!-- Contact for Large Donations -->
    <div class="bg-main text-white py-12">
      <div class="max-w-4xl mx-auto text-center">
        <h2 class="text-3xl font-bold mb-4">Planning a Major Gift?</h2>
        <p class="text-xl mb-8">
          For large donations, estate planning, or to discuss special giving
          opportunities, we'd love to speak with you personally.
        </p>
        <router-link
          to="/contact"
          class="btn-primary bg-white text-main hover:bg-gray-100"
        >
          Contact Our Finance Team
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from "vue";

export default {
  name: "Donate",
  setup() {
    const showDonationForm = ref(false);
    const donationType = ref("");
    const submitting = ref(false);

    const quickAmounts = [25, 50, 100, 250, 500, 1000];

    const donationForm = ref({
      amount: "",
      name: "",
      email: "",
      message: "",
      isAnonymous: false,
    });

    const selectDonationType = (type) => {
      donationType.value = type;
      showDonationForm.value = true;
    };

    const closeDonationForm = () => {
      showDonationForm.value = false;
      donationForm.value = {
        amount: "",
        name: "",
        email: "",
        message: "",
        isAnonymous: false,
      };
    };

    const submitDonation = async () => {
      submitting.value = true;

      try {
        // Simulate donation processing
        await new Promise((resolve) => setTimeout(resolve, 3000));

        alert(
          "Thank you for your generous donation! You will receive a confirmation email shortly.",
        );
        closeDonationForm();
      } catch (error) {
        alert("There was an error processing your donation. Please try again.");
      } finally {
        submitting.value = false;
      }
    };

    return {
      showDonationForm,
      donationType,
      submitting,
      quickAmounts,
      donationForm,
      selectDonationType,
      closeDonationForm,
      submitDonation,
    };
  },
};
</script>
