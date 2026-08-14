import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class BeliefSection extends Component {
  <template>
    <section class="section-wrapper" id="belief">
      <div class="section-header">
        <UiBadge @variant="fill">OUR BELIEF</UiBadge>
        <h2 class="section-title">
          Teknologi Seharusnya Membantu Organisasi
          <span class="brand-gradient-text">Bekerja Lebih Baik</span>
        </h2>
        <p class="section-subtitle">
          Kami percaya bahwa setiap hambatan dalam proses kerja adalah peluang
          untuk menciptakan solusi yang lebih baik. Teknologi bukan tujuan
          akhir, melainkan alat utama untuk mencapai efisiensi dan kontrol.
        </p>
      </div>

      <div class="card-grid">
        <div class="aifl-card">
          <div class="card-icon card-icon-blue">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <polygon
                points="13 2 3 14 12 14 11 22 21 10 12 10 13 2"
              ></polygon>
            </svg>
          </div>
          <h3 class="card-title">Bekerja Lebih Efisien</h3>
          <p class="card-desc">
            Mengeliminasi siklus kerja manual yang berulang dan memangkas waktu
            proses operasional harian tim.
          </p>
        </div>

        <div class="aifl-card">
          <div class="card-icon card-icon-cyan">
            <svg
              xmlns="http://www.w3.org/2000/svg"
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
                d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"
              ></path>
              <path d="M14 2v4a2 2 0 0 0 2 2h4"></path>
              <path d="M10 9H8"></path>
              <path d="M16 13H8"></path>
              <path d="M16 17H8"></path>
            </svg>
          </div>
          <h3 class="card-title">Informasi Lebih Jelas</h3>
          <p class="card-desc">
            Menyajikan data terstruktur dan transparan sehingga tidak ada lagi
            informasi penting yang tercecer.
          </p>
        </div>

        <div class="aifl-card">
          <div class="card-icon card-icon-teal">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <line x1="4" x2="4" y1="21" y2="14"></line>
              <line x1="4" x2="4" y1="10" y2="3"></line>
              <line x1="12" x2="12" y1="21" y2="12"></line>
              <line x1="12" x2="12" y1="8" y2="3"></line>
              <line x1="20" x2="20" y1="21" y2="16"></line>
              <line x1="20" x2="20" y1="12" y2="3"></line>
              <line x1="2" x2="6" y1="14" y2="14"></line>
              <line x1="10" x2="14" y1="8" y2="8"></line>
              <line x1="18" x2="22" y1="16" y2="16"></line>
            </svg>
          </div>
          <h3 class="card-title">Tingkatkan Kontrol Aktivitas</h3>
          <p class="card-desc">
            Memberikan kepastian pemantauan siapa melakukan apa dan sejauh mana
            progress pekerjaan berjalan.
          </p>
        </div>

        <div class="aifl-card">
          <div class="card-icon card-icon-purple">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <polyline points="22 7 13.5 15.5 8.5 10.5 2 17"></polyline>
              <polyline points="16 7 22 7 22 13"></polyline>
            </svg>
          </div>
          <h3 class="card-title">Keputusan Lebih Baik</h3>
          <p class="card-desc">
            Mendukung jajaran pemimpin mengambil keputusan strategis dengan
            cepat berbasis data yang akurat.
          </p>
        </div>
      </div>

      <div class="belief-quote-box">
        <h3 class="belief-quote-title">
          "Karena itu, kami tidak hanya membangun sistem. Kami membantu
          organisasi menemukan cara kerja yang lebih baik melalui teknologi."
        </h3>
      </div>
    </section>
  </template>
}
