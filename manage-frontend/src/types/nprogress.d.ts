declare module 'nprogress' {
  interface NProgressOptions {
    minimum?: number;
    template?: string;
    easing?: string;
    speed?: number;
    trickle?: boolean;
    trickleSpeed?: number;
    showSpinner?: boolean;
    parent?: string;
  }

  interface NProgress {
    configure(options: NProgressOptions): NProgress;
    start(): NProgress;
    done(force?: boolean): NProgress;
    set(n: number): NProgress;
    inc(amount?: number): NProgress;
    remove(): void;
    isStarted(): boolean;
    status: number | null;
  }

  const nprogress: NProgress;
  export = nprogress;
}

declare module 'nprogress/nprogress.css';
