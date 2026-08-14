import Component from '@glimmer/component';

export default class FooterSection extends Component {
  <template>
    <footer class="footer">
      <div class="footer-inner">
        <div class="footer-brand">
          <a href="#" class="nav-brand nav-brand-inline">
            <img src="/logo.png" alt="AIFL Services" class="brand-logo-img" />
            <div class="brand-text">
              <span class="brand-name">AIFL
                <span class="brand-name-highlight">SERVICES</span></span>
              <span class="brand-tag">Digital Solutions</span>
            </div>
          </a>
          <p>
            Membantu organisasi mendapatkan kontrol yang lebih baik atas data,
            aktivitas tim, dan proses kerja melalui solusi digital yang
            dirancang sesuai kebutuhan bisnis.
          </p>
        </div>

        <div>
          <div class="footer-col-title">Navigasi Utama</div>
          <ul class="footer-links">
            <li><a href="#about" class="footer-link">Tentang AIFL</a></li>
            <li><a href="#belief" class="footer-link">Our Belief</a></li>
            <li><a href="#problems" class="footer-link">Tantangan Operasional</a></li>
            <li><a href="#services" class="footer-link">Solusi & Layanan</a></li>
          </ul>
        </div>

        <div>
          <div class="footer-col-title">Pendekatan</div>
          <ul class="footer-links">
            <li><a href="#process" class="footer-link">Cara Kami Bekerja</a></li>
            <li><a href="#case-studies" class="footer-link">Studi Kasus</a></li>
            <li><a href="#why-us" class="footer-link">Why AIFL</a></li>
            <li><a href="#contact" class="footer-link">Process Discovery</a></li>
          </ul>
        </div>

        <div>
          <div class="footer-col-title">Solusi Digital</div>
          <ul class="footer-links">
            <li><span class="footer-link">Workflow Digitalization</span></li>
            <li><span class="footer-link">Business Process Apps</span></li>
            <li><span class="footer-link">Monitoring & Reporting</span></li>
            <li><span class="footer-link">IT Consultation</span></li>
          </ul>
        </div>
      </div>

      <div class="footer-bottom">
        <div>
          © 2026 AIFL SERVICES. All Rights Reserved. Digital Solutions for
          Better Operational Control.
        </div>
        <div class="footer-copyright-links">
          <a href="#hero" class="footer-link">Kembali ke Atas ↑</a>
        </div>
      </div>
    </footer>
  </template>
}
