import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class ProblemsSection extends Component {
  <template>
    <section class="section-wrapper" id="problems">
      <div class="section-header">
        <UiBadge @variant="danger">
          TANTANGAN OPERASIONAL
        </UiBadge>
        <h2 class="section-title">
          Tantangan Utama yang Kami Bantu <span class="brand-gradient-text">Selesaikan</span>
        </h2>
        <p class="section-subtitle">
          Banyak organisasi yang sedang berkembang menghadapi tantangan operasional mendasar berikut seiring meningkatnya skala aktivitas:
        </p>
      </div>

      <div class="card-grid">
        <div class="aifl-card problem-card">
          <div class="problem-number">01</div>
          <div class="card-icon card-icon-danger">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <ellipse cx="12" cy="5" rx="9" ry="3"></ellipse>
              <path d="M3 5v14c0 1.66 4 3 9 3s9-1.34 9-3V5"></path>
              <path d="M3 12c0 1.66 4 3 9 3s9-1.34 9-3"></path>
            </svg>
          </div>
          <h3 class="card-title">Data Sulit Dikontrol</h3>
          <p class="card-desc">
            Informasi penting masih tersebar di berbagai tempat, spreadsheet terpisah, dan pesan pribadi sehingga sulit mendapatkan gambaran yang akurat.
          </p>
        </div>

        <div class="aifl-card problem-card">
          <div class="problem-number">02</div>
          <div class="card-icon card-icon-warning">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"></path>
              <circle cx="9" cy="7" r="4"></circle>
              <path d="M22 21v-2a4 4 0 0 0-3-3.87"></path>
              <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
            </svg>
          </div>
          <h3 class="card-title">Aktivitas Tim Sulit Dipantau</h3>
          <p class="card-desc">
            Pemimpin organisasi kesulitan mengetahui siapa melakukan apa, sejauh mana progres pekerjaan, dan bagian mana yang mengalami hambatan.
          </p>
        </div>

        <div class="aifl-card problem-card">
          <div class="problem-number">03</div>
          <div class="card-icon card-icon-purple">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect width="8" height="4" x="8" y="2" rx="1" ry="1"></rect>
              <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path>
              <path d="M12 11h4"></path>
              <path d="M12 16h4"></path>
              <path d="M8 11h.01"></path>
              <path d="M8 16h.01"></path>
            </svg>
          </div>
          <h3 class="card-title">Proses Kerja Masih Manual</h3>
          <p class="card-desc">
            Banyak aktivitas masih bergantung pada spreadsheet manual, komunikasi bertumpuk, pencatatan terpisah, dan proses approval yang panjang.
          </p>
        </div>

        <div class="aifl-card problem-card">
          <div class="problem-number">04</div>
          <div class="card-icon card-icon-blue">
            <svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="10"></circle>
              <polyline points="12 6 12 12 16 14"></polyline>
            </svg>
          </div>
          <h3 class="card-title">Keputusan Butuh Waktu Lama</h3>
          <p class="card-desc">
            Kurangnya informasi terstruktur dan laporan aktual membuat manajemen membutuhkan waktu lebih lama untuk mengambil keputusan penting.
          </p>
        </div>
      </div>
    </section>
  </template>
}
