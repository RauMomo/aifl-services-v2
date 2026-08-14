import Component from '@glimmer/component';
import { pageTitle } from 'ember-page-title';
import WhatsappButton from '../components/whatsapp-button';

export default class Application extends Component {
  constructor() {
    super(...arguments);
    if (typeof window !== 'undefined') {
      let ticking = false;

      const updateParallax = () => {
        const scrollY = window.scrollY;
        const viewportHeight = window.innerHeight;
        document.documentElement.style.setProperty(
          '--scroll-y',
          `${scrollY}px`,
        );

        const sections = document.querySelectorAll(
          '.section-wrapper, .hero-section, .contact-container, .footer, .about-box, .commitment-card',
        );
        sections.forEach((sec) => {
          const rect = sec.getBoundingClientRect();
          const totalDist = viewportHeight + rect.height;
          const currentDist = viewportHeight - rect.top;
          const progress = Math.max(0, Math.min(1, currentDist / totalDist));
          const centerOffset = rect.top + rect.height / 2 - viewportHeight / 2;

          sec.style.setProperty('--sec-progress', progress.toFixed(4));
          sec.style.setProperty('--sec-offset', `${centerOffset.toFixed(1)}px`);
        });

        ticking = false;
      };

      const onScroll = () => {
        if (!ticking) {
          requestAnimationFrame(updateParallax);
          ticking = true;
        }
      };

      window.addEventListener('scroll', onScroll, { passive: true });
      window.addEventListener('resize', onScroll, { passive: true });
      updateParallax();
    }
  }

  <template>
    {{pageTitle
      "AIFL Services | Digital Solutions That Help Organizations Work Better"
    }}

    <div class="app-container">
      {{! Background Grid & Parallax Orbs }}
      <div class="bg-grid-pattern"></div>
      <div class="parallax-orb orb-1"></div>
      <div class="parallax-orb orb-2"></div>
      <div class="parallax-orb orb-3"></div>

      {{! Active Route Outlet }}
      {{outlet}}

      {{! Floating WhatsApp CTA Button }}
      <WhatsappButton />
    </div>
  </template>
}
