import { CSSProperties } from "react";

type Credit = {
  name: string | string[];
  role: string;
}

export type VideoMetadata = {
  title: string;
  description?: string;
  weight?: number; // weight of how much to show in the header.
  useboxartaspreview?: boolean;
  boxart: {
    title: string;
    caption: string;
    img: string;
    video?: string;
  };
  mobileBoxart?: {
    title: string;
    caption: string;
    img: string;
    video?: string;
  };
  links?: {
    youtube?: string;
    vimeo?: string;
  }
  credits?: Credit[];
  url: string;
  date: string;
  img?: string;
  style?: CSSProperties;
}
