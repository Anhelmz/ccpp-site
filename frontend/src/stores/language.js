import { defineStore } from "pinia";

export const useLanguageStore = defineStore("language", {
  state: () => ({
    selectedLanguage: "en",
  }),
  actions: {
    init() {
      if (typeof window !== "undefined") {
        const savedLanguage = window.localStorage.getItem("selectedLanguage");
        if (savedLanguage === "en" || savedLanguage === "kh") {
          this.selectedLanguage = savedLanguage;
        }
      }
      this.syncDocumentLanguage();
    },
    setLanguage(language) {
      if (language !== "en" && language !== "kh") return;
      if (this.selectedLanguage === language) return;

      this.selectedLanguage = language;

      if (typeof window !== "undefined") {
        window.localStorage.setItem("selectedLanguage", language);
      }

      this.syncDocumentLanguage();
    },
    syncDocumentLanguage() {
      if (typeof document === "undefined") return;
      const langCode =
        this.selectedLanguage === "kh" ? "km" : this.selectedLanguage;
      document.documentElement.setAttribute("lang", langCode);
    },
  },
});
