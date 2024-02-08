export type NewsItem = {
  title: string;
  description: string;
  url: string;
  urlToImage: string;
  publishedAt: Date;
  author: string;
  source: {
    name: string;
  }
}