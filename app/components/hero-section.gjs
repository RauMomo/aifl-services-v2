import Component from '@glimmer/component';
import { action } from '@ember/object';
import { modifier } from 'ember-modifier';
import UiButton from './ui-button';
import UiBadge from './ui-badge';

const setupScatterCanvas = modifier((canvas) => {
  const ctx = canvas.getContext('2d');
  if (!ctx) return;

  let animationFrameId;
  const parent = canvas.parentElement;
  let width = (canvas.width = parent ? parent.offsetWidth : window.innerWidth);
  let height = (canvas.height = parent ? parent.offsetHeight : 600);

  const particles = [];
  const particleCount = Math.min(Math.floor((width * height) / 9000), 90);

  const colors = ['#3B82F6', '#06B6D4', '#6366F1', '#10B981', '#60A5FA'];

  const mouse = {
    x: -1000,
    y: -1000,
    radius: 170,
  };

  const createParticles = () => {
    particles.length = 0;
    for (let i = 0; i < particleCount; i++) {
      const x = Math.random() * width;
      const y = Math.random() * height;
      particles.push({
        baseX: x,
        baseY: y,
        x: x,
        y: y,
        vx: 0,
        vy: 0,
        size: Math.random() * 2.8 + 1.2,
        color: colors[Math.floor(Math.random() * colors.length)],
        density: Math.random() * 25 + 8,
        alpha: Math.random() * 0.5 + 0.4,
      });
    }
  };

  createParticles();

  const handleResize = () => {
    if (!parent) return;
    width = canvas.width = parent.offsetWidth;
    height = canvas.height = parent.offsetHeight;
    createParticles();
  };

  window.addEventListener('resize', handleResize);

  const section = canvas.closest('.hero-section') || parent;

  const handlePointerMove = (e) => {
    if (!section) return;
    const rect = section.getBoundingClientRect();
    mouse.x = e.clientX - rect.left;
    mouse.y = e.clientY - rect.top;
  };

  const handlePointerLeave = () => {
    mouse.x = -1000;
    mouse.y = -1000;
  };

  const handleSectionClick = (e) => {
    if (!section) return;
    const rect = section.getBoundingClientRect();
    const clickX = e.clientX - rect.left;
    const clickY = e.clientY - rect.top;

    for (let i = 0; i < particles.length; i++) {
      const p = particles[i];
      const dx = p.x - clickX;
      const dy = p.y - clickY;
      const dist = Math.sqrt(dx * dx + dy * dy);
      const maxDist = 350;
      if (dist < maxDist) {
        const force = (1 - dist / maxDist) * 25;
        const angle = Math.atan2(dy, dx);
        p.vx += Math.cos(angle) * force;
        p.vy += Math.sin(angle) * force;
      }
    }
  };

  if (section) {
    section.addEventListener('mousemove', handlePointerMove);
    section.addEventListener('mouseleave', handlePointerLeave);
    section.addEventListener('click', handleSectionClick);
  }

  const render = () => {
    ctx.clearRect(0, 0, width, height);

    for (let i = 0; i < particles.length; i++) {
      const p = particles[i];

      const dx = mouse.x - p.x;
      const dy = mouse.y - p.y;
      const distance = Math.sqrt(dx * dx + dy * dy);

      if (distance < mouse.radius) {
        const force = (mouse.radius - distance) / mouse.radius;
        const directionX = dx / distance || 0;
        const directionY = dy / distance || 0;
        p.vx -= directionX * force * (p.density / 3);
        p.vy -= directionY * force * (p.density / 3);
      }

      p.vx += (p.baseX - p.x) * 0.05;
      p.vy += (p.baseY - p.y) * 0.05;

      p.vx *= 0.86;
      p.vy *= 0.86;

      p.x += p.vx;
      p.y += p.vy;

      ctx.save();
      ctx.globalAlpha = p.alpha;
      ctx.fillStyle = p.color;
      ctx.shadowColor = p.color;
      ctx.shadowBlur = 10;
      ctx.beginPath();
      ctx.arc(p.x, p.y, p.size, 0, Math.PI * 2);
      ctx.fill();
      ctx.restore();

      for (let j = i + 1; j < particles.length; j++) {
        const p2 = particles[j];
        const ldx = p.x - p2.x;
        const ldy = p.y - p2.y;
        const ldist = Math.sqrt(ldx * ldx + ldy * ldy);

        if (ldist < 95) {
          ctx.save();
          ctx.globalAlpha = (1 - ldist / 95) * 0.22;
          ctx.strokeStyle = '#3B82F6';
          ctx.lineWidth = 0.8;
          ctx.beginPath();
          ctx.moveTo(p.x, p.y);
          ctx.lineTo(p2.x, p2.y);
          ctx.stroke();
          ctx.restore();
        }
      }
    }

    animationFrameId = requestAnimationFrame(render);
  };

  render();

  return () => {
    cancelAnimationFrame(animationFrameId);
    window.removeEventListener('resize', handleResize);
    if (section) {
      section.removeEventListener('mousemove', handlePointerMove);
      section.removeEventListener('mouseleave', handlePointerLeave);
      section.removeEventListener('click', handleSectionClick);
    }
  };
});

const setupMatrixCanvas = modifier((canvas) => {
  const ctx = canvas.getContext('2d');
  if (!ctx) return;

  let animationFrameId;
  const parent = canvas.parentElement;
  let width = (canvas.width = parent ? parent.offsetWidth : 600);
  let height = (canvas.height = parent ? parent.offsetHeight : 400);

  const fontSize = 12;
  const columns = Math.floor(width / fontSize) + 1;
  const drops = Array.from({ length: columns }, () =>
    Math.floor(Math.random() * -60),
  );
  const chars =
    '0123456789ABCDEFABCDEFGHIJKLMNOPQRSTUVWXYZアァカサタナハマヤャラワガザダバパイィキシチニヒミリヰギジヂビピウゥクスツヌフムユュルグズブヅプエェケセテネヘメレヱゲゼデベペオォコソトノホモヨョロヲゴゾドボポヴッン';

  const handleResize = () => {
    if (!parent) return;
    width = canvas.width = parent.offsetWidth;
    height = canvas.height = parent.offsetHeight;
  };

  window.addEventListener('resize', handleResize);

  let lastTime = 0;
  const fps = 24;
  const interval = 1000 / fps;

  const render = (currentTime) => {
    animationFrameId = requestAnimationFrame(render);
    const delta = currentTime - lastTime;
    if (delta < interval) return;
    lastTime = currentTime - (delta % interval);

    ctx.fillStyle = 'rgba(4, 13, 8, 0.18)';
    ctx.fillRect(0, 0, width, height);

    ctx.font = `${fontSize}px "Fira Code", monospace, sans-serif`;

    for (let i = 0; i < drops.length; i++) {
      const char = chars[Math.floor(Math.random() * chars.length)];
      const x = i * fontSize;
      const y = drops[i] * fontSize;

      ctx.fillStyle = '#A7F3D0';
      ctx.shadowColor = '#10B981';
      ctx.shadowBlur = 6;
      ctx.fillText(char, x, y);

      ctx.fillStyle = '#059669';
      ctx.shadowBlur = 0;
      ctx.fillText(char, x, y - fontSize);

      if (y > height && Math.random() > 0.975) {
        drops[i] = 0;
      }
      drops[i]++;
    }
  };

  animationFrameId = requestAnimationFrame(render);

  return () => {
    cancelAnimationFrame(animationFrameId);
    window.removeEventListener('resize', handleResize);
  };
});

const setupTypingCode = modifier((container) => {
  if (!container) return;

  const CODE_LINES = [
    {
      num: '01',
      tokens: [
        { text: 'import', class: 'kw' },
        { text: ' { ' },
        { text: 'NeuralEngine', class: 'type' },
        { text: ' } ' },
        { text: 'from', class: 'kw' },
        { text: ' ' },
        { text: "'@aifl/core'", class: 'str' },
        { text: ';' },
      ],
    },
    { num: '02', tokens: [] },
    {
      num: '03',
      tokens: [{ text: '// Initializing Green Code Engine...', class: 'cmt' }],
    },
    {
      num: '04',
      tokens: [
        { text: 'const', class: 'kw' },
        { text: ' engine = ' },
        { text: 'new', class: 'kw' },
        { text: ' ' },
        { text: 'NeuralEngine', class: 'type' },
        { text: '({ mode: ' },
        { text: "'QUANTUM'", class: 'str' },
        { text: ' });' },
      ],
    },
    {
      num: '05',
      tokens: [
        { text: 'await', class: 'kw' },
        { text: ' engine.' },
        { text: 'initAutonomousLoop', class: 'fn' },
        { text: '();' },
      ],
    },
    { num: '06', tokens: [] },
    {
      num: '07',
      tokens: [
        { text: 'const', class: 'kw' },
        { text: ' stream = engine.' },
        { text: 'subscribe', class: 'fn' },
        { text: '({ throughput: ' },
        { text: "'1.4 GB/s'", class: 'str' },
        { text: ' });' },
      ],
    },
    {
      num: '08',
      tokens: [
        { text: 'stream.' },
        { text: 'on', class: 'fn' },
        { text: '(' },
        { text: "'data'", class: 'str' },
        { text: ', (' },
        { text: 'chunk', class: 'type' },
        { text: ') => {' },
      ],
    },
    {
      num: '09',
      tokens: [
        { text: '  process.stdout.' },
        { text: 'write', class: 'fn' },
        { text: '(chunk.' },
        { text: 'glow', class: 'fn' },
        { text: '());' },
      ],
    },
    { num: '10', tokens: [{ text: '});' }] },
  ];

  const LOG_ROWS = [
    { type: 'info', tag: '[SYS]', msg: 'Quantum Core #AIFL-09 ONLINE' },
    {
      type: 'success',
      tag: '[OK]',
      msg: 'Target compiled in 0.38s | GPU acceleration enabled',
    },
    {
      type: 'prompt',
      tag: '>>>',
      msg: 'aifl-daemon --execute-stream (Active 60 FPS)',
    },
  ];

  let timeoutId;
  let isDestroyed = false;

  const runTypingLoop = () => {
    if (isDestroyed) return;
    container.innerHTML = '';

    const editorBody = document.createElement('div');
    editorBody.className = 'code-editor-body';
    container.appendChild(editorBody);

    const logsPanel = document.createElement('div');
    logsPanel.className = 'terminal-logs-panel';
    logsPanel.style.opacity = '0';
    logsPanel.style.transition = 'opacity 0.4s ease';
    container.appendChild(logsPanel);

    let lineIdx = 0;
    let tokenIdx = 0;
    let charIdx = 0;

    let currentLineEl = null;
    let currentTokenEl = null;

    const startNextLine = () => {
      if (isDestroyed) return;
      if (lineIdx >= CODE_LINES.length) {
        showLogs();
        return;
      }

      const lineData = CODE_LINES[lineIdx];
      currentLineEl = document.createElement('div');
      currentLineEl.className = 'code-line';

      const lineNumEl = document.createElement('span');
      lineNumEl.className = 'line-num';
      lineNumEl.textContent = lineData.num;
      currentLineEl.appendChild(lineNumEl);

      const contentSpan = document.createElement('span');
      contentSpan.className = 'line-content';
      currentLineEl.appendChild(contentSpan);

      editorBody.appendChild(currentLineEl);

      tokenIdx = 0;
      charIdx = 0;

      if (lineData.tokens.length === 0) {
        lineIdx++;
        timeoutId = setTimeout(startNextLine, 60);
      } else {
        typeNextChar(contentSpan);
      }
    };

    const typeNextChar = (contentSpan) => {
      if (isDestroyed) return;
      const lineData = CODE_LINES[lineIdx];

      if (tokenIdx >= lineData.tokens.length) {
        lineIdx++;
        timeoutId = setTimeout(startNextLine, 80);
        return;
      }

      const currentToken = lineData.tokens[tokenIdx];
      if (charIdx === 0) {
        currentTokenEl = document.createElement('span');
        if (currentToken.class) {
          currentTokenEl.className = currentToken.class;
        }
        contentSpan.appendChild(currentTokenEl);
      }

      currentTokenEl.textContent += currentToken.text[charIdx];
      charIdx++;

      if (charIdx >= currentToken.text.length) {
        tokenIdx++;
        charIdx = 0;
      }

      timeoutId = setTimeout(() => typeNextChar(contentSpan), 20);
    };

    const showLogs = () => {
      if (isDestroyed) return;
      logsPanel.innerHTML = '';
      LOG_ROWS.forEach((log) => {
        const row = document.createElement('div');
        row.className = `log-row ${log.type}`;
        row.innerHTML = `<span class="log-tag">${log.tag}</span> <span class="log-msg">${log.msg}</span>`;
        if (log.type === 'prompt') {
          row.innerHTML += ' <span class="cursor-blk">█</span>';
        }
        logsPanel.appendChild(row);
      });
      logsPanel.style.opacity = '1';

      timeoutId = setTimeout(runTypingLoop, 3500);
    };

    startNextLine();
  };

  runTypingLoop();

  return () => {
    isDestroyed = true;
    clearTimeout(timeoutId);
  };
});

export default class HeroSection extends Component {
  setupScatterCanvas = setupScatterCanvas;
  setupMatrixCanvas = setupMatrixCanvas;
  setupTypingCode = setupTypingCode;

  @action
  handleMouseMove(event) {
    const card = event.currentTarget;
    const rect = card.getBoundingClientRect();
    const x = event.clientX - rect.left - rect.width / 2;
    const y = event.clientY - rect.top - rect.height / 2;
    const mouseX = (x / rect.width) * 14;
    const mouseY = -(y / rect.height) * 14;
    card.style.setProperty('--mouse-x', `${mouseX}deg`);
    card.style.setProperty('--mouse-y', `${mouseY}deg`);
  }

  @action
  handleMouseLeave(event) {
    const card = event.currentTarget;
    card.style.setProperty('--mouse-x', '0deg');
    card.style.setProperty('--mouse-y', '0deg');
  }

  <template>
    <section class="hero-section" id="hero">
      <canvas class="hero-scatter-canvas" {{this.setupScatterCanvas}}></canvas>

      {{! Animated Right Orbital System }}
      <div class="hero-orbital-system" aria-hidden="true">
        <div class="orbital-ring-layer orbital-layer-outer">
          <svg
            class="orbital-svg"
            viewBox="0 0 800 800"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle cx="400" cy="400" r="385" class="orbit-faint-line" />
            <circle cx="400" cy="400" r="360" class="orbit-outer-segmented" />
            <circle cx="124" cy="168" r="11" class="orbit-node-primary" />
            <circle cx="462" cy="754" r="10" class="orbit-node-primary" />
            <circle cx="676" cy="168" r="6" class="orbit-node-secondary" />
          </svg>
        </div>

        <div class="orbital-ring-layer orbital-layer-mid">
          <svg
            class="orbital-svg"
            viewBox="0 0 800 800"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle cx="400" cy="400" r="270" class="orbit-mid-dashed" />
            <circle cx="166" cy="535" r="7" class="orbit-node-teal" />
            <circle cx="660" cy="330" r="5" class="orbit-node-teal" />
          </svg>
        </div>

        <div class="orbital-ring-layer orbital-layer-inner">
          <svg
            class="orbital-svg"
            viewBox="0 0 800 800"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <circle cx="400" cy="400" r="175" class="orbit-inner-arc" />
            <circle cx="400" cy="400" r="85" class="orbit-core-circle" />
            <line x1="655" y1="145" x2="675" y2="125" class="orbit-axis-tick" />
            <line x1="125" y1="675" x2="145" y2="655" class="orbit-axis-tick" />
          </svg>
        </div>
      </div>

      <div class="hero-content">
        <UiBadge @showDot={{true}} @variant="fill">
          AIFL SERVICES — Your One Stop Solution for your IT needs
        </UiBadge>

        <h1 class="hero-main-title">
          Digital Solutions That Help Organizations
          <span class="brand-gradient-text">Work Better</span>
        </h1>

        <p class="hero-description">
          Membantu organisasi mendapatkan kontrol yang lebih baik atas data,
          aktivitas tim, dan proses kerja melalui solusi digital yang dirancang
          sesuai kebutuhan bisnis.
        </p>

        <div class="hero-cta">
          <UiButton @href="#contact" @variant="primary">
            Diskusi Kebutuhan Bisnis
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
          </UiButton>
          <UiButton @href="#about" @variant="secondary">
            Pelajari Pendekatan Kami
          </UiButton>
        </div>
      </div>

      <div class="hero-visual-wrapper">
        {{! template-lint-disable no-invalid-interactive }}
        <div
          class="hero-parallax-card laptop-showcase-card"
          {{on "mousemove" this.handleMouseMove}}
          {{on "mouseleave" this.handleMouseLeave}}
        >
          {{! Laptop Graphic Container }}
          <div class="laptop-container">
            {{! Laptop Screen Lid }}
            <div class="laptop-screen">
              <div class="laptop-camera"></div>
              <div class="laptop-screen-glass">
                <canvas
                  class="laptop-matrix-canvas"
                  {{this.setupMatrixCanvas}}
                ></canvas>
                <div class="terminal-scanlines"></div>

                {{! Terminal Window Frame }}
                <div class="laptop-terminal-window">
                  <div class="terminal-header">
                    <div class="terminal-dots">
                      <span class="dot dot-close"></span>
                      <span class="dot dot-min"></span>
                      <span class="dot dot-max"></span>
                    </div>
                    <div class="terminal-tabs">
                      <div class="terminal-tab active">
                        <span class="tab-icon">⚡</span>
                        green_code_engine.rs
                      </div>
                    </div>
                    <div class="terminal-status">
                      <span class="status-pulse"></span>
                      <span class="status-text">LIVE STREAM</span>
                    </div>
                  </div>

                  <div
                    class="terminal-content-view matrix-view fade-in"
                    {{this.setupTypingCode}}
                  >
                  </div>

                  <div class="screen-glow-overlay"></div>
                </div>
              </div>
            </div>

            {{! Laptop Base Keyboard }}
            <div class="laptop-base">
              <div class="laptop-notch"></div>
              <div class="laptop-reflection"></div>
            </div>
          </div>
        </div>

        {{! Floating Badges for Depth Parallax Layer }}
        <div class="floating-badge float-top-right">
          <div>
            <div class="badge-title">AIFL Software Design</div>
            <div class="badge-desc">Mockup UI/UX Kustom & Modern</div>
          </div>
        </div>

        <div class="floating-badge float-bottom-left">
          <div>
            <div class="badge-title">Responsif & Intuitif</div>
            <div class="badge-desc">Dirancang khusus untuk tim Anda</div>
          </div>
        </div>
      </div>
    </section>
  </template>
}
