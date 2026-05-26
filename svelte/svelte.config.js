import adapter from '@sveltejs/adapter-static';

export default {
  kit: {
    adapter: adapter({
      pages: '../golang/public',
      assets: '../golang/public',
      fallback: 'index.html',
      precompress: false,
      strict: true
    })
  }
};
