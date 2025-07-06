export interface GenericServerHealth<D> {
  server_name: string; // Name of the server
  build_date: D | 'unknown';
  version: string; // Version of the server
  commit: string; // Commit hash of the server
  branch: string; // Branch of the server 
  environment: string
}

export type ServerHealth = GenericServerHealth<string>;
