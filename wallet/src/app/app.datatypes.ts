export class Wallet {
  label: string;
}

export class WalletsNewRequest {
  label: string;
  seed: string;
  aCount: number;
  encrypted: boolean;
  password: string;
}
