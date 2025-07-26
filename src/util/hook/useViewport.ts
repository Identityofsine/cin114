"use client";
import { useEffect, useState } from "react";

export default function useViewport() {
  const [viewport, setViewport] = useState({
    width: 1920, // Default fallback width
    height: 1080, // Default fallback height
  });

  useEffect(() => {
    const handleResize = () => {
      setViewport({
        width: window.innerWidth,
        height: window.innerHeight,
      });
    };

    // Initial size - only run on client
    handleResize();

    window.addEventListener('resize', handleResize);

    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return viewport;
}
