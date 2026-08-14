import Component from '@glimmer/component';

export default class UiButton extends Component {
  get variantClass() {
    const variant = this.args.variant || 'fill';
    return `btn-${variant}`;
  }

  get sizeClass() {
    const size = this.args.size;
    return size ? `btn-${size}` : '';
  }

  get buttonType() {
    return this.args.type || 'button';
  }

  <template>
    {{#if @href}}
      <a
        href={{@href}}
        class="btn {{this.variantClass}} {{this.sizeClass}}"
        ...attributes
      >
        {{yield}}
      </a>
    {{else}}
      <button
        type={{this.buttonType}}
        disabled={{@disabled}}
        class="btn {{this.variantClass}} {{this.sizeClass}}"
        ...attributes
      >
        {{yield}}
      </button>
    {{/if}}
  </template>
}
