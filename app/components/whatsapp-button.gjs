import Component from '@glimmer/component';

export default class WhatsappButton extends Component {
  // Default WhatsApp Consultation Link & Pre-filled Message
  whatsappUrl =
    'https://wa.me/6281234567890?text=Halo%20AIFL%20Services%2C%20saya%20ingin%20berkonsultasi%20mengenai%20solusi%20digital%20dan%20otomatisasi%20operasional';

  <template>
    <a
      href={{this.whatsappUrl}}
      target="_blank"
      rel="noopener noreferrer"
      class="whatsapp-fab"
      aria-label="Konsultasi via WhatsApp bersama AIFL Services"
    >
      <span class="whatsapp-pulse"></span>
      <span class="whatsapp-tooltip">
        Konsultasi via WhatsApp
      </span>
      <div class="whatsapp-icon-box">
        <svg
          width="28"
          height="28"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path
            d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"
          ></path>
        </svg>
      </div>
    </a>
  </template>
}
