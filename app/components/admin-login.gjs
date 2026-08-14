import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { service } from '@ember/service';
import { on } from '@ember/modifier';
import UiBadge from './ui-badge';
import UiButton from './ui-button';

export default class AdminLoginComponent extends Component {
  @service auth;
  @service router;

  @tracked usernameOrEmail = '';
  @tracked password = '';
  @tracked showPassword = false;
  @tracked rememberMe = true;

  @action
  togglePasswordVisibility() {
    this.showPassword = !this.showPassword;
  }

  @action
  handleInputUsername(event) {
    this.usernameOrEmail = event.target.value;
  }

  @action
  handleInputPassword(event) {
    this.password = event.target.value;
  }

  @action
  async handleSubmit(event) {
    event.preventDefault();
    if (!this.usernameOrEmail || !this.password) return;

    const success = await this.auth.login(this.usernameOrEmail, this.password);
    if (success) {
      this.router.transitionTo('portal-access.dashboard');
    }
  }

  <template>
    <div class="admin-portal-container fade-in">
      <div class="admin-portal-card">
        {{! Left Block: IT Cyber Control Room Image Showcase }}
        <div class="admin-portal-left">
          <img
            src="/admin_portal_hero.png"
            alt="AIFL IT Infrastructure Control Center"
            class="admin-hero-image"
          />
          <div class="admin-left-overlay">
            <div class="admin-status-badge">
              <span class="status-dot green"></span>
              SYSTEM SECURITY LEVEL 4
            </div>

            <div class="admin-left-text-group">
              <h2 class="admin-left-title">AIFL Core Control Portal</h2>
              <p class="admin-left-subtitle">
                Interface khusus administrator untuk mengelola workflow, sistem audit, dan pengawasan operasional digital.
              </p>

              <div class="admin-metrics-pill">
                <span class="lock-icon">🔒</span>
                <span>ENCRYPTION: AES-256-GCM | PROTOBUF ACTIVE</span>
              </div>
            </div>
          </div>
        </div>

        {{! Right Block: Sleek Dark Glassmorphic Login Form }}
        <div class="admin-portal-right">
          <div class="admin-form-header">
            <div class="admin-brand-logo">
              <UiBadge @variant="fill">AKSES TERPROTEKSI</UiBadge>
            </div>
            <h1 class="admin-form-title">Masuk Portal Admin</h1>
            <p class="admin-form-subtitle">
              Silakan masukkan kredensial administrator Anda untuk melanjutkan.
            </p>
          </div>

          {{#if this.auth.error}}
            <div class="admin-error-box">
              <span class="error-icon">⚠️</span>
              <span>{{this.auth.error}}</span>
            </div>
          {{/if}}

          <form {{on "submit" this.handleSubmit}} class="admin-login-form">
            <div class="form-group">
              <label for="admin-user-input" class="form-label">Username atau Email</label>
              <div class="input-wrapper">
                <span class="input-icon">👤</span>
                <input
                  id="admin-user-input"
                  type="text"
                  class="admin-input"
                  placeholder="admin"
                  value={{this.usernameOrEmail}}
                  {{on "input" this.handleInputUsername}}
                  required
                />
              </div>
            </div>

            <div class="form-group">
              <label for="admin-pass-input" class="form-label">Kata Sandi</label>
              <div class="input-wrapper">
                <span class="input-icon">🔑</span>
                <input
                  id="admin-pass-input"
                  type={{if this.showPassword "text" "password"}}
                  class="admin-input"
                  placeholder="••••••••"
                  value={{this.password}}
                  {{on "input" this.handleInputPassword}}
                  required
                />
                <button
                  type="button"
                  class="toggle-pass-btn"
                  {{on "click" this.togglePasswordVisibility}}
                  aria-label="Toggle visibility"
                >
                  {{if this.showPassword "👁️" "🙈"}}
                </button>
              </div>
            </div>

            <div class="form-options">
              <label class="remember-label">
                <input type="checkbox" checked={{this.rememberMe}} />
                <span>Ingat Sesi Ini</span>
              </label>
            </div>

            <div class="form-submit-row">
              <UiButton
                @type="submit"
                @variant="primary"
                class="admin-submit-btn"
                disabled={{this.auth.isLoading}}
              >
                {{#if this.auth.isLoading}}
                  <span class="spinner-sm"></span>
                  Memverifikasi...
                {{else}}
                  Masuk ke Portal
                  <svg
                    width="18"
                    height="18"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2.5"
                  ><line x1="5" y1="12" x2="19" y2="12"></line><polyline
                      points="12 5 19 12 12 19"
                    ></polyline></svg>
                {{/if}}
              </UiButton>
            </div>
          </form>
        </div>
      </div>
    </div>
  </template>
}
