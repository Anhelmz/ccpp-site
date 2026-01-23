<template>
  <div>
    <!-- Hero Section -->
    <div class="py-6 md:py-8">
      <div class="max-w-4xl mx-auto text-center px-4 sm:px-6 lg:px-8">
        <h1 class="text-4xl md:text-5xl font-bold mb-2">Resources</h1>
      </div>
    </div>

    <!-- Content Section -->
    <div class="bg-primary py-12 md:py-16">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
        <p class="text-lg text-gray-700 text-center mb-12 max-w-2xl mx-auto">
          Tools and materials for your spiritual journey.
        </p>

        <div class="prose prose-lg max-w-none">
          <p class="text-lg text-gray-700 mb-6">
            We've compiled helpful resources to support you on your spiritual
            journey. Whether you're looking for Bible study materials, sermon
            recordings, prayer guides, or information about our programs, we're
            here to help.
          </p>
          <p class="text-lg text-gray-700 mb-4">
            Our available resources include:
          </p>
          <ul class="space-y-3 text-gray-700 mb-6">
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>Bible Study Materials:</strong> Study guides,
                commentaries, weekly study guides, Bible reading plans, and
                devotional materials</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>Sermon Recordings:</strong> Audio recordings, video
                messages, and message transcripts of our weekly teachings</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>Prayer Resources:</strong> Prayer guides, intercession
                lists, and prayer request forms</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>Library Resources:</strong> Christian books, reference
                materials, and educational resources available in our church
                library</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>Family Resources:</strong> Family devotionals,
                parenting resources, and children's materials</span
              >
            </li>
            <li class="flex items-start">
              <span class="text-main mr-2">•</span>
              <span
                ><strong>New Believer Resources:</strong> Getting started
                guides, foundation classes, and discipleship materials for those
                new to the faith</span
              >
            </li>
          </ul>
          <p class="text-lg text-gray-700 mb-8">
            Our resources are designed to help you grow in your relationship
            with God, understand Scripture better, and connect with our church
            community. If you need specific materials or have questions, please
            contact us.
          </p>
        </div>

        <div class="mt-12 mb-12">
          <h2 class="text-2xl font-semibold text-gray-900 text-center mb-6">
            Follow Us on Social Media
          </h2>
          <div class="flex flex-wrap justify-center gap-4">
            <a
              href="https://www.facebook.com/calvarychapelpp"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center justify-center px-6 py-3 bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-700 transition-colors shadow-md"
            >
              <font-awesome-icon :icon="['fab', 'facebook']" class="w-5 h-5 mr-2" />
              Facebook
            </a>
            <a
              href="https://www.youtube.com"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center justify-center px-6 py-3 bg-red-600 text-white font-semibold rounded-lg hover:bg-red-700 transition-colors shadow-md"
            >
              <font-awesome-icon :icon="['fab', 'youtube']" class="w-5 h-5 mr-2" />
              YouTube
            </a>
          </div>
        </div>

        <div class="flex flex-wrap justify-center gap-4 mt-12">
          <router-link
            to="/contact"
            class="inline-flex items-center justify-center px-6 py-3 text-base bg-brand-orange text-white font-semibold rounded-lg hover:bg-brand-orange/90 hover:shadow-xl transform hover:-translate-y-0.5 transition-all duration-300 shadow-lg"
          >
            Contact Us
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Resources",
};
</script>
