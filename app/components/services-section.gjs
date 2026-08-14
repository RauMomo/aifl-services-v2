import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import { on } from '@ember/modifier';
import { fn } from '@ember/helper';
import UiBadge from './ui-badge';
import UiButton from './ui-button';

export default class ServicesSection extends Component {
  @tracked activeServiceIndex = 0;

  servicesData = [
    {
      id: 'workflow',
      title: 'Workflow Digitalization',
      badge: 'ALUR KERJA OTOMATIS',
      subtitle: 'Mengubah proses manual menjadi workflow digital terstruktur',
      desc: 'Kami mendigitalisasi alur kerja manual organisasi agar setiap tahap pekerjaan memiliki kejelasan status, penanggung jawab, serta otomatisasi alur persetujuan tanpa hambatan.',
      examples: ['Proses Approval Otomatis', 'Monitoring Pekerjaan', 'Pengelolaan Tim', 'Tracking Progress Real-time'],
      iconType: 'workflow'
    },
    {
      id: 'custom-app',
      title: 'Custom Business Application',
      badge: 'APLIKASI PRESISI',
      subtitle: 'Membangun aplikasi khusus yang disesuaikan dengan kebutuhan bisnis',
      desc: 'Bukan software generik yang kaku. Kami merancang dan membangun aplikasi custom yang presisi mengikuti kebutuhan operasional dan cara kerja nyata bisnis Anda.',
      examples: ['Sistem Operational Control', 'Portal Manajemen Terpadu', 'Aplikasi Kasus Khusus', 'Integrasi Multi-Platform'],
      iconType: 'app'
    },
    {
      id: 'monitoring',
      title: 'Monitoring & Reporting System',
      badge: 'VISIBILITAS REAL-TIME',
      subtitle: 'Membantu organisasi mendapatkan informasi yang mudah dipantau',
      desc: 'Menghadirkan visibilitas penuh melalui dashboard intuitif dan laporan berkala untuk memantau performa operasional dan indikator bisnis utama secara akurat.',
      examples: ['Executive Dashboard', 'Laporan Automatic Generation', 'Tracking KPIs & SLA', 'Analisis Performa Bisnis'],
      iconType: 'monitoring'
    },
    {
      id: 'consultation',
      title: 'IT Consultation & Discovery',
      badge: 'RANCANGAN STRATEGIS',
      subtitle: 'Membantu organisasi memetakan kebutuhan & arsitektur solusi digital',
      desc: 'Melakukan analisis mendalam terhadap alur bisnis Anda untuk menemukan titik hambatan dan memetakan rancangan arsitektur sistem digital sebelum tahap pengembangan.',
      examples: ['Problem Discovery', 'Business Process Mapping', 'Requirement Engineering', 'Arsitektur Sistem'],
      iconType: 'consultation'
    }
  ];

  get servicesWithState() {
    return this.servicesData.map((service, index) => ({
      ...service,
      index,
      isActive: index === this.activeServiceIndex
    }));
  }

  get currentService() {
    return this.servicesData[this.activeServiceIndex] || this.servicesData[0];
  }

  @action
  selectService(index) {
    this.activeServiceIndex = index;
  }

  <template>
    <section class="section-wrapper" id="services">
      <div class="section-header">
        <UiBadge @variant="fill">SOLUSI KAMI</UiBadge>
        <h2 class="section-title">
          Solusi Digital yang Sesuai dengan <span class="brand-gradient-text">Proses Kerja Anda</span>
        </h2>
        <p class="section-subtitle">
          AIFL membantu organisasi membangun solusi digital yang dirancang khusus berdasarkan alur dan kebutuhan operasional nyata.
        </p>
      </div>

      {{!-- Inline 50% / 50% Equal Split Layout Container --}}
      <div class="services-split-layout">
        {{!-- Left Column (50% Width) - Service Selector Boxes (Title Only) --}}
        <div class="services-selector-column">
          {{#each this.servicesWithState as |service|}}
            {{! template-lint-disable no-invalid-interactive }}
            <div 
              class="service-tab-card {{if service.isActive 'active'}}"
              {{on "click" (fn this.selectService service.index)}}
            >
              <div class="service-card-left">
                <div class="service-icon-box">
                  {{#if (eq service.iconType 'workflow')}}
                    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="7" rx="1"></rect><rect x="14" y="3" width="7" height="7" rx="1"></rect><rect x="14" y="14" width="7" height="7" rx="1"></rect><rect x="3" y="14" width="7" height="7" rx="1"></rect><path d="M10 6h4"></path><path d="M17 10v4"></path><path d="M10 17h4"></path></svg>
                  {{else if (eq service.iconType 'app')}}
                    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="20" height="14" x="2" y="3" rx="2"></rect><line x1="8" y1="21" x2="16" y2="21"></line><line x1="12" y1="17" x2="12" y2="21"></line></svg>
                  {{else if (eq service.iconType 'monitoring')}}
                    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="18" y1="20" x2="18" y2="10"></line><line x1="12" y1="20" x2="12" y2="4"></line><line x1="6" y1="20" x2="6" y2="14"></line></svg>
                  {{else}}
                    <svg width="22" height="22" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 14c.2-1 .7-1.7 1.5-2.5 1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5.7.7 1.3 1.5 1.5 2.5"></path><path d="M9 18h6"></path><path d="M10 22h4"></path></svg>
                  {{/if}}
                </div>
                <h3 class="service-card-title">{{service.title}}</h3>
              </div>

              <div class="service-card-indicator">
                {{#if service.isActive}}
                  <span class="active-dot"></span>
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><polyline points="9 18 15 12 9 6"></polyline></svg>
                {{else}}
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" opacity="0.3"><polyline points="9 18 15 12 9 6"></polyline></svg>
                {{/if}}
              </div>
            </div>
          {{/each}}
        </div>

        {{!-- Right Column (50% Width) - Detailed Selected Service Showcase --}}
        <div class="services-detail-column">
          <div class="services-showcase-panel">
            <div class="services-showcase-info fade-in">
              <div class="services-detail-badge">
                <UiBadge @variant="fill">{{this.currentService.badge}}</UiBadge>
              </div>
              <h3 class="services-detail-title">
                {{this.currentService.title}}
              </h3>
              <div class="services-detail-subtitle">
                {{this.currentService.subtitle}}
              </div>
              <p class="services-detail-desc">
                {{this.currentService.desc}}
              </p>

              <div class="services-chips-group">
                {{#each this.currentService.examples as |item|}}
                  <span class="service-chip">
                    <span class="chip-icon">✓</span> {{item}}
                  </span>
                {{/each}}
              </div>

              <div>
                <UiButton @href="#contact" @variant="primary">
                  Diskusi Solusi Ini
                  <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5"><line x1="5" y1="12" x2="19" y2="12"></line><polyline points="12 5 19 12 12 19"></polyline></svg>
                </UiButton>
              </div>
            </div>

            {{!-- Interactive Visual Preview Widget --}}
            <div class="services-showcase-visual fade-in">
              {{#if (eq this.currentService.id 'workflow')}}
                <div class="service-preview-widget workflow-widget">
                  <div class="widget-header">
                    <span class="widget-dot green"></span>
                    <span class="widget-title">Workflow Execution Pipeline</span>
                    <span class="widget-status">ACTIVE PIPELINE</span>
                  </div>
                  <div class="pipeline-steps">
                    <div class="pipeline-step completed">
                      <div class="step-num">01</div>
                      <div class="step-label">Pengajuan</div>
                      <div class="step-check">✓</div>
                    </div>
                    <div class="pipeline-arrow">➔</div>
                    <div class="pipeline-step completed">
                      <div class="step-num">02</div>
                      <div class="step-label">Validasi</div>
                      <div class="step-check">✓</div>
                    </div>
                    <div class="pipeline-arrow">➔</div>
                    <div class="pipeline-step active">
                      <div class="step-num">03</div>
                      <div class="step-label">Approval</div>
                      <div class="status-pulse"></div>
                    </div>
                  </div>
                  <div class="widget-footer-bar">
                    <div class="bar-progress" style="width: 75%;"></div>
                  </div>
                </div>
              {{else if (eq this.currentService.id 'custom-app')}}
                <div class="service-preview-widget app-widget">
                  <div class="widget-header">
                    <span class="widget-dot blue"></span>
                    <span class="widget-title">AIFL Custom Operational Control</span>
                    <span class="widget-status">SYSTEM READY</span>
                  </div>
                  <div class="widget-app-body">
                    <div class="app-mini-card">
                      <div class="mini-label">TOTAL TASK OPERASIONAL</div>
                      <div class="mini-value">1,420 Active</div>
                    </div>
                    <div class="app-mini-card">
                      <div class="mini-label">OTOMATISASI WORKFLOW</div>
                      <div class="mini-value cyan">99.4%</div>
                    </div>
                  </div>
                  <div class="widget-list">
                    <div class="widget-list-item"><span class="item-dot"></span> Modul Penjadwalan Synchronized</div>
                    <div class="widget-list-item"><span class="item-dot"></span> Tracking Database Real-time</div>
                  </div>
                </div>
              {{else if (eq this.currentService.id 'monitoring')}}
                <div class="service-preview-widget monitoring-widget">
                  <div class="widget-header">
                    <span class="widget-dot teal"></span>
                    <span class="widget-title">Live Performance Dashboard</span>
                    <span class="widget-status">REAL-TIME DATA</span>
                  </div>
                  <div class="chart-bars">
                    <div class="chart-col"><div class="chart-bar" style="height: 65%;"></div><span>SEN</span></div>
                    <div class="chart-col"><div class="chart-bar" style="height: 85%;"></div><span>SEL</span></div>
                    <div class="chart-col"><div class="chart-bar active" style="height: 100%;"></div><span>RAB</span></div>
                    <div class="chart-col"><div class="chart-bar" style="height: 70%;"></div><span>KAM</span></div>
                    <div class="chart-col"><div class="chart-bar" style="height: 90%;"></div><span>JUM</span></div>
                  </div>
                  <div class="widget-stats-row">
                    <span>Data Sync: <strong>100%</strong></span>
                    <span>SLA Compliance: <strong>99.9%</strong></span>
                  </div>
                </div>
              {{else}}
                <div class="service-preview-widget discovery-widget">
                  <div class="widget-header">
                    <span class="widget-dot amber"></span>
                    <span class="widget-title">Process Discovery & Blueprint</span>
                    <span class="widget-status">DIAGNOSTIC ROADMAP</span>
                  </div>
                  <div class="discovery-nodes">
                    <div class="disc-node"><span class="disc-icon">🔍</span> 1. Problem Discovery</div>
                    <div class="disc-connector"></div>
                    <div class="disc-node"><span class="disc-icon">🗺️</span> 2. Business Process Mapping</div>
                    <div class="disc-connector"></div>
                    <div class="disc-node active"><span class="disc-icon">⚡</span> 3. Solution Architecture</div>
                  </div>
                </div>
              {{/if}}
            </div>
          </div>
        </div>
      </div>
    </section>
  </template>
}

