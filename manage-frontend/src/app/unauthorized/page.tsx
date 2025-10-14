import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { AlertTriangle, ArrowLeft } from "lucide-react";

/**
 * 无权限访问页面
 */
export default function UnauthorizedPage() {
  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center px-4">
      <Card className="w-full max-w-md">
        <CardHeader className="text-center">
          <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-red-100">
            <AlertTriangle className="h-6 w-6 text-red-600" />
          </div>
          <CardTitle className="mt-4 text-xl font-semibold text-gray-900">
            访问被拒绝
          </CardTitle>
        </CardHeader>
        <CardContent className="text-center space-y-4">
          <p className="text-gray-600">
            抱歉，您没有权限访问此页面。请联系管理员获取相应权限。
          </p>
          <div className="flex flex-col space-y-2">
            <Link href="/dashboard">
              <Button className="w-full">
                <ArrowLeft className="mr-2 h-4 w-4" />
                返回首页
              </Button>
            </Link>
            <Link href="/login">
              <Button variant="outline" className="w-full">
                重新登录
              </Button>
            </Link>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
