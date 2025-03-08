import createMiddleware from 'next-intl/middleware';
import { NextRequest } from 'next/server';
import { uniqid } from './utils/utils';

export async function middleware(request: NextRequest) {

  const response = createMiddleware({
    // A list of all locales that are supported
    locales: ['en', 'fa'],
    localeDetection: false,
    // Used when no locale matches
    defaultLocale: 'en',
    localePrefix: 'as-needed'
  })(request);

  if (!request.cookies.has('session')) {
    const uniq: any = uniqid()
    response.cookies.set({
      name: 'session',
      value: uniq,
      path: '/',
    })
  }
  
  const url = new URL(request.url);
  const origin = url.origin;
  const pathname = url.pathname;

  response.headers.set('x-url', request.url);
  response.headers.set('x-origin', origin);
  response.headers.set('x-pathname', pathname);

  return response
}

export const config = {
  // Match only internationalized pathnames
  matcher: [
    '/((?!api|_next/static|_next/image|favicon.ico|apple-touch-icon.png|favicon.svg|logo.svg|images/books|icons|manifest|flags|images|assets|img|js).*)'
  ]
};
