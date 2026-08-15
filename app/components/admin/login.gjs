import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { service } from '@ember/service';
import { on } from '@ember/modifier';
import UiBadge from '../ui-badge';
import UiButton from '../ui-button';

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
            <div class="admin-left-text-group">
              <h2 class="admin-left-title">AIFL Core Control Portal</h2>
              <p class="admin-left-subtitle">
                Interface khusus administrator untuk mengelola workflow, sistem audit, dan pengawasan operasional digital.
              </p>

            </div>
          </div>
        </div>

        {{! Right Block: Sleek Dark Glassmorphic Login Form }}
        <div class="admin-portal-right">
          <div class="admin-form-header">
            <div class="admin-brand-logo">
              <UiBadge @variant="fill">AIFL Services (Khusus Admin)</UiBadge>
            </div>
            <h1 class="admin-form-title">Masuk Dashboard Admin</h1>
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
                <span class="input-icon">
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"></path>
                    <circle cx="12" cy="7" r="4"></circle>
                  </svg>
                </span>
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
                <span class="input-icon">
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect width="18" height="11" x="3" y="11" rx="2" ry="2"></rect>
                    <path d="M7 11V7a5 5 0 0 1 10 0v4"></path>
                  </svg>
                </span>
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
                  {{#if this.showPassword}}
                    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path>
                      <circle cx="12" cy="12" r="3"></circle>
                    </svg>
                  {{else}}
                    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                      <path d="M9.88 9.88a3 3 0 1 0 4.24 4.24"></path>
                      <path d="M10.73 5.08A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"></path>
                      <path d="M6.61 6.61A13.52 13.52 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61"></path>
                      <line x1="2" x2="22" y1="2" y2="22"></line>
                    </svg>
                  {{/if}}
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
