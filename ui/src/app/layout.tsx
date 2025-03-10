import type { Metadata } from "next";
import { getTranslations } from "next-intl/server";

export const metadata: Metadata = {
  title: "Free Subtitle Api",
  description: "Free Subtitle Api",
  keywords: "Download Subtitle,Download Free Subtitle"
};

export default async function RootLayout({
  children,
  params
}: Readonly<{
  children: React.ReactNode;
  params: Promise<{ locale: string }>
}>) {

  const { locale } = await params;

  const t = await getTranslations()

  return (
    <html lang={locale}>
      <body dir={t('dir')}>
        {children}
      </body>
    </html>
  );
}
