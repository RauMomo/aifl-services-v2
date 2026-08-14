import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { on } from '@ember/modifier';
import UiButton from './ui-button';
import UiBadge from './ui-badge';

const BACK4APP_APP_ID = import.meta.env?.VITE_BACK4APP_APP_ID || 'LaSGnbphXz9sI1gjSpqPrFsIm07XJ1b1wJbCcpAD';
const BACK4APP_CLIENT_KEY = import.meta.env?.VITE_BACK4APP_CLIENT_KEY || '4VAYoABttwcicWUL6aX76G7IsvWr8XbsnNnjWTAo';
const BACK4APP_SERVER_URL = import.meta.env?.VITE_BACK4APP_SERVER_URL || 'https://parseapi.back4app.com';

export default class ContactSection extends Component {
  @tracked name = '';
  @tracked organization = '';
  @tracked email = '';
  @tracked message = '';
  @tracked isSubmitting = false;
  @tracked isSubmitted = false;
  @tracked submitError = '';

  @action
  handleNameInput(event) {
    this.name = event.target.value;
  }

  @action
  handleOrgInput(event) {
    this.organization = event.target.value;
  }

  @action
  handleEmailInput(event) {
    this.email = event.target.value;
  }

  @action
  handleMessageInput(event) {
    this.message = event.target.value;
  }

  @action
  async handleSubmit(event) {
    event.preventDefault();
    if (!this.name || !this.email || this.isSubmitting) return;

    this.isSubmitting = true;
    this.submitError = '';

    try {
      const response = await fetch(`${BACK4APP_SERVER_URL}/classes/Contact`, {
        method: 'POST',
        headers: {
          'X-Parse-Application-Id': BACK4APP_APP_ID,
          'X-Parse-Client-Key': BACK4APP_CLIENT_KEY,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: this.name,
          organization: this.organization,
          email: this.email,
          message: this.message,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.error || `HTTP ${response.status}`);
      }

      this.isSubmitted = true;
    } catch (err) {
      console.error('Failed to submit contact request to Back4App:', err);
      this.submitError = 'Terjadi kesalahan saat mengirim pesan. Silakan periksa koneksi atau coba lagi nanti.';
    } finally {
      this.isSubmitting = false;
    }
  }

  @action
  resetForm() {
    this.name = '';
    this.organization = '';
    this.email = '';
    this.message = '';
    this.isSubmitting = false;
    this.isSubmitted = false;
    this.submitError = '';
  }

  <template>
    <section class="section-wrapper" id="contact">
      <div class="contact-container">
        <div>
          <UiBadge @variant="fill">MULAI DISKUSI</UiBadge>
          <h2 class="contact-title">
            Let's Build <span class="brand-gradient-text">Better Ways of Working</span>
          </h2>
          <p class="contact-subtitle">
            AIFL Services — Digital Solutions for Better Operational Control
          </p>
          <p class="contact-desc">
            Siap mengubah alur kerja manual menjadi sistem digital yang teratur dan terkontrol? Diskusikan tantangan operasional organisasi Anda bersama tim ahli kami.
          </p>

          <div class="contact-info-list">
            <div class="contact-info-item">
              <div class="contact-info-icon">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect width="20" height="16" x="2" y="4" rx="2"></rect>
                  <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"></path>
                </svg>
              </div>
              <div>
                <div class="contact-info-label">Email Konsultasi</div>
                <div class="contact-info-value">hello@aiflservices.com</div>
              </div>
            </div>

            <div class="contact-info-item">
              <div class="contact-info-icon">
                <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M16 20V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"></path>
                  <rect width="20" height="14" x="2" y="6" rx="2"></rect>
                </svg>
              </div>
              <div>
                <div class="contact-info-label">Layanan Utamakan</div>
                <div class="contact-info-value">Process Discovery & Custom App Dev</div>
              </div>
            </div>
          </div>
        </div>

        <div>
          {{#if this.isSubmitted}}
            <div class="contact-success-box">
              <div class="contact-success-icon">🎉</div>
              <h3 class="contact-success-title">
                Pesan Konsultasi Terkirim!
              </h3>
              <p class="contact-success-desc">
                Terima kasih {{this.name}} dari {{this.organization}}. Tim AIFL akan meninjau kebutuhan operasional Anda dan menghubungi dalam 1x24 jam.
              </p>
              <UiButton @variant="muted" {{on "click" this.resetForm}}>
                Kirim Pesan Lain
              </UiButton>
            </div>
          {{else}}
            <form {{on "submit" this.handleSubmit}} class="contact-form">
              {{#if this.submitError}}
                <div style="background-color: rgba(239, 68, 68, 0.1); border: 1px solid rgba(239, 68, 68, 0.3); color: #f87171; padding: 0.75rem 1rem; border-radius: 0.5rem; font-size: 0.875rem;">
                  {{this.submitError}}
                </div>
              {{/if}}

              <div>
                <label for="contact-name" class="form-label">
                  Nama Lengkap *
                </label>
                <input 
                  id="contact-name"
                  type="text" 
                  required 
                  disabled={{this.isSubmitting}}
                  placeholder="Nama Anda" 
                  value={{this.name}}
                  {{on "input" this.handleNameInput}}
                  class="form-input"
                />
              </div>

              <div>
                <label for="contact-org" class="form-label">
                  Nama Organisasi / Perusahaan
                </label>
                <input 
                  id="contact-org"
                  type="text" 
                  disabled={{this.isSubmitting}}
                  placeholder="Perusahaan / Instansi Anda" 
                  value={{this.organization}}
                  {{on "input" this.handleOrgInput}}
                  class="form-input"
                />
              </div>

              <div>
                <label for="contact-email" class="form-label">
                  Email Bisnis *
                </label>
                <input 
                  id="contact-email"
                  type="email" 
                  required 
                  disabled={{this.isSubmitting}}
                  placeholder="email@organisasi.com" 
                  value={{this.email}}
                  {{on "input" this.handleEmailInput}}
                  class="form-input"
                />
              </div>

              <div>
                <label for="contact-message" class="form-label">
                  Tantangan Operasional yang Dihadapi
                </label>
                <textarea 
                  id="contact-message"
                  rows="3" 
                  disabled={{this.isSubmitting}}
                  placeholder="Ceritakan proses manual atau hambatan yang ingin didigitalisasi..." 
                  value={{this.message}}
                  {{on "input" this.handleMessageInput}}
                  class="form-textarea"
                ></textarea>
              </div>

              <UiButton @type="submit" @variant="fill" class="contact-submit-btn" disabled={{this.isSubmitting}}>
                {{#if this.isSubmitting}}
                  Mengirim Pesan...
                {{else}}
                  Jadwalkan Sesi Consultation & Process Discovery
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="5" y1="12" x2="19" y2="12"></line><polyline points="12 5 19 12 12 19"></polyline></svg>
                {{/if}}
              </UiButton>
            </form>
          {{/if}}
        </div>
      </div>
    </section>
  </template>
}
