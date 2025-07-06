export const StorageLks = {
  AUTH: 'authToken'
}

type StorageKeys = keyof typeof StorageLks;

type ValidStorageKeys = typeof StorageLks[StorageKeys];

export class Storage {

  safeGetItem(key: ValidStorageKeys): string | null {
    try {
      // Check if we're in a browser environment
      if (typeof window === 'undefined') {
        return null;
      }
      return localStorage.getItem(key);
    } catch (e) {
      console.error(e);
      return null;
    }
  }

  safeSetItem(key: ValidStorageKeys, value: string): void {
    try {
      // Check if we're in a browser environment
      if (typeof window === 'undefined') {
        return;
      }
      localStorage.setItem(key, value);
    } catch (e) {
      console.error(e);
    }
  }

  safeRemoveItem(key: ValidStorageKeys): void {
    try {
      // Check if we're in a browser environment
      if (typeof window === 'undefined') {
        return;
      }
      localStorage.removeItem(key);
    } catch (e) {
      console.error(e);
    }
  }
}
