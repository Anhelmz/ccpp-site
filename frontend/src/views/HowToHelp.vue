<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">How to Help</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Join us in serving God's Kingdom through prayer, giving, and service.
        </p>

        <div class="prose prose-lg max-w-none">
          <h2 class="text-2xl font-semibold text-gray-900 mb-4">Prayer</h2>
          <p class="text-lg text-gray-700 mb-6">
            Prayer is the foundation of everything we do. We believe in the
            power of prayer to transform lives, heal communities, and bring
            people closer to God. You can help through prayer by joining our
            weekly prayer meetings, signing up for our prayer chain, committing
            to pray for our missionaries and outreach programs, and
            participating in our monthly prayer and fasting days.
          </p>

          <h2 class="text-2xl font-semibold text-gray-900 mb-4">Donation</h2>
          <p class="text-lg text-gray-700 mb-6">
            Your generous donations help us continue our mission of spreading
            the Gospel, serving our community, and supporting those in need.
            Every contribution, no matter the size, makes a significant
            difference. We have several funds including General Fund, Mission
            Fund, Building Fund, and Benevolence Fund.
          </p>

          <h2 class="text-2xl font-semibold text-gray-900 mb-4">Service</h2>
          <p class="text-lg text-gray-700 mb-6">
            God has given each of us unique gifts and talents to serve His
            church and community. We invite you to use your abilities to make a
            difference. Ministry opportunities include Worship Team, Children's
            Ministry, Youth Ministry, Hospitality Team, Outreach & Evangelism,
            Facilities & Maintenance, Communications & Media, and Administrative
            Support.
          </p>

          <h2 class="text-2xl font-semibold text-gray-900 mb-4">Needs List</h2>
          <p class="text-lg text-gray-700 mb-4">
            We are grateful for the support of our community. Current ministry
            needs include:
          </p>
          <ul class="space-y-3 text-gray-700 mb-8">
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                >Audio/Visual equipment, musical instruments, chairs and tables,
                Bibles and study materials, office supplies</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                >Food and hygiene items, children's ministry supplies, school
                supplies, transportation support, volunteers</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                >Monthly financial support for missionaries, care packages,
                technology for communication, translation services</span
              >
            </li>
          </ul>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-lg hover:bg-brand-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Contact Us
          </router-link>
          <router-link
            to="/donate"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-white text-main font-semibold rounded-lg hover:bg-primary hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Donate Now
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "HowToHelp",
};
</script>
