import { LoadingSpinner } from "@/components/common/LoadingSpinner";

/**
 * 首页组件
 * 由于设置了自动重定向，用户不会看到此页面内容
 * 未登录用户 → 重定向到 /login
 * 已登录用户 → 重定向到 /dashboard
 */
export default function HomePage() {
  return (
    <div className="flex min-h-screen items-center justify-center">
      <LoadingSpinner text="正在加载..." />
    </div>
  );
}
