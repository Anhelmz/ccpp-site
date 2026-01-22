import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import About from '@/views/About.vue'
import Leadership from '@/views/Leadership.vue'
import WhatWeBelieve from '@/views/WhatWeBelieve.vue'
import CalvaryChapel from '@/views/CalvaryChapel.vue'
import Cambodia from '@/views/Cambodia.vue'
import Ministries from '@/views/Ministries.vue'
import YouthMinistry from '@/views/YouthMinistry.vue'
import GraceChurch from '@/views/GraceChurch.vue'
import Outreaches from '@/views/Outreaches.vue'
import Calendar from '@/views/Calendar.vue'
import UpcomingEvents from '@/views/UpcomingEvents.vue'
import Resources from '@/views/Resources.vue'
import Contact from '@/views/Contact.vue'
import Users from '@/views/Users.vue'
import Donate from '@/views/Donate.vue'
import HowToHelp from '@/views/HowToHelp.vue'
import AdminLogin from '@/views/admin/Login.vue'
import AdminDashboard from '@/views/admin/Dashboard.vue'
import GalleryManagement from '@/views/admin/GalleryManagement.vue'
import EventManagement from '@/views/admin/EventManagement.vue'
import ContentManagement from '@/views/admin/ContentManagement.vue'
import ContactMessages from '@/views/admin/ContactMessages.vue'
import Settings from '@/views/admin/Settings.vue'
import VideoManagement from '@/views/admin/VideoManagement.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'About',
    component: About
  },
  {
    path: '/about/leadership',
    name: 'Leadership',
    component: Leadership
  },
  {
    path: '/about/what-we-believe',
    name: 'WhatWeBelieve',
    component: WhatWeBelieve
  },
  {
    path: '/about/calvary-chapel',
    name: 'CalvaryChapel',
    component: CalvaryChapel
  },
  {
    path: '/about/cambodia',
    name: 'Cambodia',
    component: Cambodia
  },
  {
    path: '/ministries',
    name: 'Ministries',
    component: Ministries
  },
  {
    path: '/ministries/youth-ministry',
    name: 'YouthMinistry',
    component: YouthMinistry
  },
  {
    path: '/ministries/grace-church',
    name: 'GraceChurch',
    component: GraceChurch
  },
  {
    path: '/ministries/outreaches',
    name: 'Outreaches',
    component: Outreaches
  },
  {
    path: '/events/calendar',
    name: 'Calendar',
    component: Calendar
  },
  {
    path: '/events/upcoming-events',
    name: 'UpcomingEvents',
    component: UpcomingEvents
  },
  {
    path: '/resources',
    name: 'Resources',
    component: Resources
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact
  },
  {
    path: '/users',
    name: 'Users',
    component: Users,
    meta: { requiresAdmin: true }
  },
  {
    path: '/donate',
    name: 'Donate',
    component: Donate
  },
  {
    path: '/how-to-help',
    name: 'HowToHelp',
    component: HowToHelp
  },
  {
    path: '/admin/login',
    name: 'AdminLogin',
    component: AdminLogin
  },
  {
    path: '/admin',
    redirect: '/admin/dashboard'
  },
  {
    path: '/admin/dashboard',
    name: 'AdminDashboard',
    component: AdminDashboard,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/gallery',
    name: 'GalleryManagement',
    component: GalleryManagement,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/gallery/:category',
    name: 'GalleryCategory',
    component: GalleryManagement,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/events',
    name: 'EventManagement',
    component: EventManagement,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/calendar',
    redirect: '/admin/events'
  },
  {
    path: '/admin/videos',
    name: 'VideoManagement',
    component: VideoManagement,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/content',
    name: 'ContentManagement',
    component: ContentManagement,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/contacts',
    name: 'ContactMessages',
    component: ContactMessages,
    meta: { requiresAdmin: true }
  },
  {
    path: '/admin/settings',
    name: 'Settings',
    component: Settings,
    meta: { requiresAdmin: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Temporary bypass to allow admin views without login while testing.
const BYPASS_ADMIN_GUARD = false

router.beforeEach((to, from, next) => {
  if (BYPASS_ADMIN_GUARD) {
    return next()
  }

  const requiresAdmin = to.matched.some(record => record.meta?.requiresAdmin)
  const token = localStorage.getItem('authToken') || localStorage.getItem('auth_token')
  const userRaw = localStorage.getItem('authUser')
  const user = userRaw ? JSON.parse(userRaw) : null

  if (requiresAdmin) {
    if (!token || !user || user.role !== 'admin') {
      return next({
        name: 'AdminLogin',
        query: { redirect: to.fullPath }
      })
    }
  }

  if (to.name === 'AdminLogin' && token && user?.role === 'admin') {
    return next({ name: 'AdminDashboard' })
  }

  next()
})

export default router
