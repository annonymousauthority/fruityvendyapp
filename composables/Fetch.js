// fetch.js
import { ref } from "vue";

export async function fetchFruity(url) {
  const data = ref(null);
  const error = ref(null);

  try {
    const response = await fetch(url);
    const data = await response.json();
    return data;
  } catch (error) {
    console.error(error);
  }
}
