import type { Metadata } from "next";
import { BrandSettings } from "@/brand.settings";
import './thank-you.scss';

export const metadata: Metadata = {
  title: BrandSettings.brandName.displayShort + ' - Thank You',
  description: 'Thank you for your purchase! Check your email for your tickets.',
  twitter: {
    title: BrandSettings.brandName.displayShort + ' - Thank You',
    description: 'Thank you for your purchase! Check your email for your tickets.',
    card: 'summary',
    site: '@cin114',
  },
  openGraph: {
    title: BrandSettings.brandName.displayShort + ' - Thank You',
    description: 'Thank you for your purchase! Check your email for your tickets.',
    type: 'website',
    locale: 'en_US',
  },
};

export default function ThankYou() {
  return (
    <section className="thank-you">
      {/* Hero Section */}
      <div className="thank-you__hero">
        <div className="thank-you__hero-content">
          <div className="thank-you__icon">
            <svg viewBox="0 0 100 100" className="thank-you__check">
              <circle cx="50" cy="50" r="45" className="thank-you__check-circle" />
              <path d="M25 50 L40 65 L75 30" className="thank-you__check-mark" />
            </svg>
          </div>
          <h1 className="thank-you__title">Thank You!</h1>
          <p className="thank-you__subtitle">Your purchase was successful</p>
        </div>
      </div>

      {/* Content Sections */}
      <div className="thank-you__content">
        {/* Confirmation Section */}
        <div className="thank-you__section">
          <h2 className="thank-you__section-title">Confirmation Details</h2>
          <div className="thank-you__info-grid">
            <div className="thank-you__info-item">
              <h3>Email Confirmation</h3>
              <p>We&apos;ve sent your ticket confirmation to your email address. Please check your inbox and spam folder.</p>
            </div>
            <div className="thank-you__info-item">
              <h3>Event Details</h3>
              <p>All screening information, location details, and special instructions are included in your confirmation email.</p>
            </div>
          </div>
        </div>

        <div className="divider" />

        {/* Next Steps Section */}
        <div className="thank-you__section">
          <h2 className="thank-you__section-title">What&apos;s Next?</h2>
          <div className="thank-you__steps">
            <div className="thank-you__step">
              <div className="thank-you__step-number">1</div>
              <div className="thank-you__step-content">
                <h3>Check Your Email</h3>
                <p>Look for your ticket confirmation with all the event details</p>
              </div>
            </div>
            <div className="thank-you__step">
              <div className="thank-you__step-number">2</div>
              <div className="thank-you__step-content">
                <h3>Save the Date</h3>
                <p>Add the screening to your calendar so you don&apos;t miss it</p>
              </div>
            </div>
            <div className="thank-you__step">
              <div className="thank-you__step-number">3</div>
              <div className="thank-you__step-content">
                <h3>Join Us</h3>
                <p>Arrive early to get the best seats and enjoy the experience</p>
              </div>
            </div>
          </div>
        </div>

        <div className="divider" />

        {/* Footer Message */}
        <div className="thank-you__footer-section">
          <div className="thank-you__brand">
            <img src="/home/logo.svg" alt="CIN114" className="thank-you__logo" />
          </div>
          <p className="thank-you__closing">
            We can&apos;t wait to share this cinematic experience with you.
            See you at the screening!
          </p>
          <div className="thank-you__contact">
            <p>Questions? Contact us at <a href={`mailto:${BrandSettings.contact.email}`}>{BrandSettings.contact.email}</a></p>
          </div>
        </div>
      </div>
    </section>
  );
} 
