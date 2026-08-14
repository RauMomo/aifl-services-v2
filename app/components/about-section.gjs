import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class AboutSection extends Component {
  <template>
    <section class="section-wrapper" id="about">
      <div class="about-box">
        <div>
          <UiBadge @variant="muted">TENTANG AIFL</UiBadge>
          <h2 class="section-title">
            Memahami Cara Kerja Organisasi
            <span class="brand-gradient-text">Sebelum Teknologi</span>
          </h2>

          <p class="about-desc">
            Setiap organisasi memiliki cara kerja yang unik. Namun seiring
            pertumbuhan bisnis, proses yang sebelumnya sederhana dapat menjadi
            semakin kompleks.
          </p>

          <div class="about-bullet-list">
            <div class="about-bullet-item">
              <span class="bullet-dot">•</span>
              <span class="bullet-text">Data mulai tersebar di berbagai tempat
                tanpa struktur baku.</span>
            </div>
            <div class="about-bullet-item">
              <span class="bullet-dot">•</span>
              <span class="bullet-text">Aktivitas tim semakin sulit dipantau
                secara real-time.</span>
            </div>
            <div class="about-bullet-item">
              <span class="bullet-dot">•</span>
              <span class="bullet-text">Koordinasi antar pihak membutuhkan lebih
                banyak waktu dan energi.</span>
            </div>
            <div class="about-bullet-item">
              <span class="bullet-dot">•</span>
              <span class="bullet-text">Pengambilan keputusan terhambat karena
                informasi tidak tersedia cepat.</span>
            </div>
          </div>
        </div>

        <div>
          <div class="highlight-quote-box">
            <p class="quote-text">
              Kami tidak memulai dari teknologi. Kami memulai dari memahami
              bagaimana sebuah organisasi bekerja.
            </p>
            <div class="quote-footer">
              Dengan memahami proses, hambatan, dan kebutuhan pengguna, kami
              membantu mengubah proses kerja manual menjadi solusi digital yang
              terkontrol, mudah dipantau, dan siap mendukung pertumbuhan bisnis.
            </div>
          </div>
        </div>
      </div>
    </section>
  </template>
}
