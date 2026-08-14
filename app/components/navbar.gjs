import Component from '@glimmer/component';
import UiButton from './ui-button';

export default class Navbar extends Component {
  <template>
    <header class="navbar">
      <div class="nav-inner">
        <a href="#" class="nav-brand">
          <img src="/logo.png" alt="AIFL Services" class="brand-logo-img" />
          <div class="brand-text">
            <span class="brand-name">AIFL
              <span class="brand-name-highlight">SERVICES</span></span>
            <span class="brand-tag">Digital Solutions</span>
          </div>
        </a>

        <nav>
          <ul class="nav-links">
            <li><a href="#about" class="nav-link">Tentang</a></li>
            <li><a href="#belief" class="nav-link">Prinsip</a></li>
            <li><a href="#problems" class="nav-link">Tantangan</a></li>
            <li><a href="#services" class="nav-link">Solusi Kami</a></li>
            <li><a href="#process" class="nav-link">Cara Kerja</a></li>
            <li><a href="#case-studies" class="nav-link">Studi Kasus</a></li>
            <li><a href="#why-us" class="nav-link">Why AIFL</a></li>
          </ul>
        </nav>

        <div class="nav-actions">
          <UiButton @href="#contact" @variant="fill">
            <svg
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2.5"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path
                d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"
              ></path>
            </svg>
            Konsultasi Gratis
          </UiButton>
        </div>
      </div>
    </header>
  </template>
}
