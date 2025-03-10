import type { Metadata } from "next";
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

  return (
    <html lang={locale}>
      <body>
        {children}
      </body>
    </html>
  );
}
