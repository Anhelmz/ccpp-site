const STORAGE_KEY = "ccpp_videos";

const readVideos = () => {
  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) return [];
    const parsed = JSON.parse(raw);
    if (!Array.isArray(parsed)) return [];
    return parsed;
  } catch {
    return [];
  }
};

const writeVideos = (videos) => {
  try {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(videos));
  } catch {
    /* ignore storage errors */
  }
};

export const videoStore = {
  getAll() {
    return readVideos().sort((a, b) => (b.createdAt || 0) - (a.createdAt || 0));
  },
  add(video) {
    const list = readVideos();
    const now = Date.now();
    const id = crypto.randomUUID
      ? crypto.randomUUID()
      : `vid-${now}-${list.length + 1}`;
    list.unshift({
      id,
      title: video.title,
      description: video.description,
      youtubeId: video.youtubeId,
      youtubeUrl: video.youtubeUrl,
      createdAt: now,
    });
    writeVideos(list);
    return this.getAll();
  },
  update(id, video) {
    const list = readVideos().map((item) =>
      item.id === id
        ? {
            ...item,
            title: video.title,
            description: video.description,
            youtubeId: video.youtubeId,
            youtubeUrl: video.youtubeUrl,
          }
        : item,
    );
    writeVideos(list);
    return this.getAll();
  },
  remove(id) {
    const list = readVideos().filter((item) => item.id !== id);
    writeVideos(list);
    return this.getAll();
  },
};
