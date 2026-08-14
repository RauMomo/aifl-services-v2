import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class CaseStudiesSection extends Component {
  <template>
    <section class="section-wrapper" id="case-studies">
      <div class="section-header">
        <UiBadge @variant="fill">STUDI KASUS & IMPLIKASI NYATA</UiBadge>
        <h2 class="section-title">
          Kisah Sukses Implementasi <span class="brand-gradient-text">Solusi Digital AIFL</span>
        </h2>
        <p class="section-subtitle">
          Bagaimana AIFL membantu organisasi bertransformasi dari alur kerja manual menjadi sistem yang terkontrol dan terstruktur.
        </p>
      </div>

      {{!-- Case Study 1: Drive Master --}}
      <div class="case-study-card">
        <div class="case-study-info">
          <div class="case-study-tag">
            CASE STUDY 01
          </div>
          <h3 class="case-study-title">
            Drive Master
          </h3>
          <div class="case-study-subtitle">
            Digitalizing Course Management Process
          </div>

          <div class="case-study-block">
            <div class="case-study-label-danger">
              Challenge:
            </div>
            <p class="case-study-text">
              Sebelum menggunakan solusi digital, proses utama seperti pendaftaran, pengelolaan data, dan penjadwalan masih dilakukan secara manual. Hal tersebut membuat operasional sulit dikelola ketika aktivitas bisnis meningkat pesat.
            </p>
          </div>

          <div>
            <div class="case-study-label-blue">
              Solution:
            </div>
            <p class="case-study-text">
              AIFL membangun sistem digital terintegrasi untuk mengelola data peserta, mengatur jadwal otomatis, mendukung operasional harian, dan meningkatkan keteraturan workflow.
            </p>
          </div>
        </div>

        <div class="case-study-preview">
          <div class="case-study-preview-box">
            <div class="case-study-preview-title">
              Real Operational Impact
            </div>
            
            <div class="impact-list">
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Proses operasional menjadi jauh lebih terstruktur.</span>
              </div>
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Informasi pendaftaran & jadwal mudah dikelola.</span>
              </div>
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Aktivitas operasional terkontrol 100%.</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      {{!-- Case Study 2: Internal Audit / AuditSphere --}}
      <div class="case-study-card">
        <div class="case-study-info">
          <div class="case-study-tag">
            CASE STUDY 02
          </div>
          <h3 class="case-study-title">
            Internal Audit / AuditSphere
          </h3>
          <div class="case-study-subtitle">
            Digitalizing Audit Monitoring Process
          </div>

          <div class="case-study-block">
            <div class="case-study-label-danger">
              Challenge:
            </div>
            <p class="case-study-text">
              Proses audit dan mitigasi risiko sebelumnya membutuhkan banyak pemantauan manual. Tantangan utama terletak pada tracking pekerjaan, monitoring progress, dan pencatatan aktivitas yang tersebar.
            </p>
          </div>

          <div>
            <div class="case-study-label-blue">
              Solution:
            </div>
            <p class="case-study-text">
              AIFL membangun sistem digital terpusat untuk monitoring aktivitas audit, pengelolaan timeline, tracking progress pekerjaan, dan dokumentasi proses secara sistematis.
            </p>
          </div>
        </div>

        <div class="case-study-preview">
          <div class="case-study-preview-box">
            <div class="case-study-preview-title">
              Real Operational Impact
            </div>
            
            <div class="impact-list">
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Kontrol proses audit menjadi lebih presisi.</span>
              </div>
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Monitoring progress terlihat jelas & aktual.</span>
              </div>
              <div class="impact-item">
                <span class="impact-check">✓</span>
                <span class="impact-text">Dokumentasi aktivitas audit terstruktur.</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </template>
}
