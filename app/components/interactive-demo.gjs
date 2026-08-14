import Component from '@glimmer/component';
import { tracked } from '@glimmer/tracking';
import { action } from '@ember/object';
import UiButton from './ui-button';

export default class InteractiveDemo extends Component {
  @tracked count = 42;
  @tracked activeTab = 'counter';
  @tracked textInput = 'Hello, Ember 5!';
  @tracked isGlowEnabled = true;

  @action
  increment() {
    this.count++;
  }

  @action
  decrement() {
    this.count--;
  }

  @action
  reset() {
    this.count = 0;
  }

  @action
  setTab(tab) {
    this.activeTab = tab;
  }

  @action
  updateInput(event) {
    this.textInput = event.target.value;
  }

  @action
  toggleGlow() {
    this.isGlowEnabled = !this.isGlowEnabled;
  }

  <template>
    <section class="demo-section" id="interactive-demo">
      <div class="section-header">
        <div class="section-tag">Interactive Glimmer Demo</div>
        <h2 class="section-title">Experience Ember Reactivity Live</h2>
        <p class="section-subtitle">
          See Glimmer tracked properties and single-way data binding in action.
        </p>
      </div>

      <div class="demo-container">
        <div class="demo-tabs">
          <button
            type="button"
            class="demo-tab-btn {{if (eq this.activeTab 'counter') 'active'}}"
            {{on "click" (fn this.setTab 'counter')}}
          >
            Tracked Counter
          </button>
          <button
            type="button"
            class="demo-tab-btn {{if (eq this.activeTab 'text') 'active'}}"
            {{on "click" (fn this.setTab 'text')}}
          >
            Live Two-Way Input
          </button>
          <button
            type="button"
            class="demo-tab-btn {{if (eq this.activeTab 'style') 'active'}}"
            {{on "click" (fn this.setTab 'style')}}
          >
            Dynamic Class Modifier
          </button>
        </div>

        <div class="demo-content">
          {{#if (eq this.activeTab 'counter')}}
            <div class="demo-interactive-box">
              <span class="code-title">@tracked count</span>
              <div class="counter-display">{{this.count}}</div>
              <div class="counter-controls">
                <UiButton @variant="muted" @size="sm" {{on "click" this.decrement}}>- Decrease</UiButton>
                <UiButton @variant="fill" @size="sm" {{on "click" this.increment}}>+ Increase</UiButton>
                <UiButton @variant="muted" @size="sm" {{on "click" this.reset}}>Reset</UiButton>
              </div>
            </div>
            <div>
              <h4 class="demo-section-heading">Glimmer Engine Tracking</h4>
              <p class="demo-section-body demo-section-body--spaced">
                Glimmer uses fine-grained autotracking. When <code>@tracked count</code> updates, Ember re-renders only the exact DOM node bound to the count, giving ultra-fast 60fps performance without virtual DOM diff overhead.
              </p>
            </div>
          {{else if (eq this.activeTab 'text')}}
            <div class="demo-interactive-box">
              <span class="code-title">Live Preview Output</span>
              <div class="demo-live-preview">
                "{{this.textInput}}"
              </div>
              <input
                type="text"
                value={{this.textInput}}
                {{on "input" this.updateInput}}
                placeholder="Type something here..."
                class="demo-text-input"
                aria-label="Live preview input text"
              />
            </div>
            <div>
              <h4 class="demo-section-heading">Seamless Event Binding</h4>
              <p class="demo-section-body">
                With <code>&#123;&#123;on "input" this.updateInput&#125;&#125;</code> modifiers and clean class methods, standard DOM events hook directly into your tracked state with zero boilerplate.
              </p>
            </div>
          {{else if (eq this.activeTab 'style')}}
            <div class="demo-interactive-box {{if this.isGlowEnabled 'demo-interactive-box--glow'}}">
              <span class="code-title">Card Glow Status</span>
              <div class="demo-glow-label {{if this.isGlowEnabled 'demo-glow-label--active'}}">
                {{if this.isGlowEnabled "Glow Accent Active" "Glow Accent Disabled"}}
              </div>
              <UiButton @variant="fill" @size="sm" {{on "click" this.toggleGlow}}>
                Toggle Glow State
              </UiButton>
            </div>
            <div>
              <h4 class="demo-section-heading">Template Logic Helpers</h4>
              <p class="demo-section-body">
                Declarative inline conditions like <code>&#123;&#123;if this.isGlowEnabled ...&#125;&#125;</code> keep your templates expressive, readable, and perfectly synchronized with application logic.
              </p>
            </div>
          {{/if}}
        </div>
      </div>
    </section>
  </template>
}
