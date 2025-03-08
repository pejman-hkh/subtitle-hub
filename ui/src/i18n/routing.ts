import {defineRouting} from 'next-intl/routing';
 
export const routing = defineRouting({
  locales: ['en', 'fa'],
  localeDetection: false,
  // Used when no locale matches
  defaultLocale: 'en',
  localePrefix: 'as-needed'
});