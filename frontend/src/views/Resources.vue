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
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M24 12.073c0-6.627-5.373-12-12-12s-12 5.373-12 12c0 5.99 4.388 10.954 10.125 11.854v-8.385H7.078v-3.47h3.047V9.43c0-3.007 1.792-4.669 4.533-4.669 1.312 0 2.686.235 2.686.235v2.953H15.83c-1.491 0-1.956.925-1.956 1.874v2.25h3.328l-.532 3.47h-2.796v8.385C19.612 23.027 24 18.062 24 12.073z"
                />
              </svg>
              Facebook
            </a>
            <a
              href="https://www.youtube.com"
              target="_blank"
              rel="noopener noreferrer"
              class="inline-flex items-center justify-center px-6 py-3 bg-red-600 text-white font-semibold rounded-lg hover:bg-red-700 transition-colors shadow-md"
            >
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24">
                <path
                  d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z"
                />
              </svg>
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
