import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class WhyAiflSection extends Component {
  <template>
    <section class="section-wrapper" id="why-us">
      <div class="section-header">
        <UiBadge @variant="muted">WHY AIFL</UiBadge>
        <h2 class="section-title">
          Alasan Organisasi Memilih <span class="brand-gradient-text">AIFL Services</span>
        </h2>
        <p class="section-subtitle">
          Kami berfokus pada nilai dan dampak operasional nyata, bukan sekadar menyerahkan baris kode program.
        </p>
      </div>

      <div class="card-grid grid-3-col">
        <div class="aifl-card aifl-card-flex">
          <div class="card-icon card-icon-blue-soft">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <circle cx="12" cy="12" r="6"></circle>
              <circle cx="12" cy="12" r="2"></circle>
            </svg>
          </div>
          <h3 class="card-title">Business Understanding Before Technology</h3>
          <p class="card-desc">
            Kami percaya solusi digital yang baik tidak pernah dimulai dari pamer teknologi. Solusi yang tepat selalu dimulai dari memahami masalah dan alur kerja nyata organisasi.
          </p>
        </div>

        <div class="aifl-card aifl-card-flex">
          <div class="card-icon card-icon-cyan-soft">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polygon points="12 2 2 7 12 12 22 7 12 2"></polygon>
              <polyline points="2 17 12 22 22 17"></polyline>
              <polyline points="2 12 12 17 22 12"></polyline>
            </svg>
          </div>
          <h3 class="card-title">Solution Tailored to Your Process</h3>
          <p class="card-desc">
            Setiap organisasi memiliki cara kerja yang khas. Kami membangun aplikasi yang fleksibel mengikuti proses bisnis Anda, bukan memaksa tim Anda menyesuaikan diri dengan sistem yang kaku.
          </p>
        </div>

        <div class="aifl-card aifl-card-flex">
          <div class="card-icon card-icon-teal-soft">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.71.79-1.81.2-2.55L4.5 16.5z"></path>
              <path d="M12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-3.05 11a22.35 22.35 0 0 1-3.95 2z"></path>
              <path d="M9 12H4s.5-3.5 3-5l2 2"></path>
              <path d="M15 12v5s-3.5-.5-5-3l2-2"></path>
            </svg>
          </div>
          <h3 class="card-title">Focus on Real Operational Impact</h3>
          <p class="card-desc">
            Kami tidak sekadar mengejar aplikasi yang bisa dijalankan. Fokus utama kami adalah memastikan solusi tersebut benar-benar membantu organisasi bekerja lebih efisien dan terkontrol.
          </p>
        </div>
      </div>
    </section>
  </template>
}
