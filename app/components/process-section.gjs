import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class ProcessSection extends Component {
  <template>
    <section class="section-wrapper" id="process">
      <div class="section-header">
        <UiBadge @variant="muted">CARA KAMI BEKERJA</UiBadge>
        <h2 class="section-title">
          5 Tahap Metodologi <span class="brand-gradient-text">Pengembangan Solusi</span>
        </h2>
        <p class="section-subtitle">
          Pendekatan terstruktur kami memastikan solusi digital yang dibangun benar-benar menjawab kebutuhan operasional dan dapat diadopsi dengan sukses.
        </p>
      </div>

      <div class="process-timeline">
        <div class="process-step-card">
          <div class="step-badge">1</div>
          <h3 class="process-step-title">
            Understand
          </h3>
          <p class="process-step-desc">
            Memahami proses berjalan, hambatan yang terjadi, kebutuhan pengguna, dan tujuan bisnis organisasi secara mendalam.
          </p>
        </div>

        <div class="process-step-card">
          <div class="step-badge">2</div>
          <h3 class="process-step-title">
            Analyze
          </h3>
          <p class="process-step-desc">
            Mengidentifikasi alur yang dapat diperbaiki, aktivitas yang didigitalisasi, serta data yang wajib dikontrol.
          </p>
        </div>

        <div class="process-step-card">
          <div class="step-badge">3</div>
          <h3 class="process-step-title">
            Design
          </h3>
          <p class="process-step-desc">
            Merancang solusi digital yang spesifik. Bukan solusi generik, melainkan sesuai kondisi nyata bisnis Anda.
          </p>
        </div>

        <div class="process-step-card">
          <div class="step-badge">4</div>
          <h3 class="process-step-title">
            Develop
          </h3>
          <p class="process-step-desc">
            Membangun sistem mengutamakan kemudahan pengguna, stabilitas tinggi, dan kesiapan skalabilitas masa depan.
          </p>
        </div>

        <div class="process-step-card">
          <div class="step-badge">5</div>
          <h3 class="process-step-title">
            Support
          </h3>
          <p class="process-step-desc">
            Pendampingan penuh memastikan solusi benar-benar aktif digunakan dan memberikan dampak operasional positif.
          </p>
        </div>
      </div>
    </section>
  </template>
}
