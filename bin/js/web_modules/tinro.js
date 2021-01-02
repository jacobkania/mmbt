import { G as noop, s as safe_not_equal, n as getContext, z as setContext, J as onMount, K as tick, S as SvelteComponent, i as init, u as update_slot, t as transition_in, j as transition_out, p as empty, d as insert, w as group_outros, y as check_outros, k as detach, c as create_slot } from './common/index-766c86e7.js';

const subscriber_queue = [];
/**
 * Create a `Writable` store that allows both updating and reading by subscription.
 * @param {*=}value initial value
 * @param {StartStopNotifier=}start start and stop notifications for subscriptions
 */
function writable(value, start = noop) {
    let stop;
    const subscribers = [];
    function set(new_value) {
        if (safe_not_equal(value, new_value)) {
            value = new_value;
            if (stop) { // store is ready
                const run_queue = !subscriber_queue.length;
                for (let i = 0; i < subscribers.length; i += 1) {
                    const s = subscribers[i];
                    s[1]();
                    subscriber_queue.push(s, value);
                }
                if (run_queue) {
                    for (let i = 0; i < subscriber_queue.length; i += 2) {
                        subscriber_queue[i][0](subscriber_queue[i + 1]);
                    }
                    subscriber_queue.length = 0;
                }
            }
        }
    }
    function update(fn) {
        set(fn(value));
    }
    function subscribe(run, invalidate = noop) {
        const subscriber = [run, invalidate];
        subscribers.push(subscriber);
        if (subscribers.length === 1) {
            stop = start(set) || noop;
        }
        run(value);
        return () => {
            const index = subscribers.indexOf(subscriber);
            if (index !== -1) {
                subscribers.splice(index, 1);
            }
            if (subscribers.length === 0) {
                stop();
                stop = null;
            }
        };
    }
    return { set, update, subscribe };
}

function m(e){let c=getContext("tinro");c&&(c.exact||c.fallback)&&R(`${e.fallback?"<Route fallback>":`<Route path="${e.path}">`}  can't be inside ${c.fallback?"<Route fallback>":`<Route path="${c.path||"/"}"> with exact path`}`);let a=e.fallback?"fallbacks":"childs",t={un:null,exact:!1,pattern:"",params:{},parent:c,fallback:e.fallback,redirect:e.redirect,firstmatch:e.firstmatch,matched:!1,childs:new Set,activeChilds:new Set,fallbacks:new Set,makePattern(r){t.exact=!r.endsWith("/*"),t.pattern=f(`${t.parent&&t.parent.pattern||""}${r}`);},register:()=>t.parent?(t.parent[a].add(t),()=>{t.parent[a].delete(t),t.un&&t.un();}):void 0,show:()=>{e.onShow(),!t.fallback&&t.parent&&t.parent.activeChilds.add(t);},hide:()=>{e.onHide(),!t.fallback&&t.parent&&t.parent.activeChilds.delete(t);},match:async r=>{t.matched=!1;let n=h(t.pattern,r);if(n&&t.redirect&&(!t.exact||t.exact&&n.exact))return o.goto(v(r,t.parent.pattern,t.redirect));if(n&&!t.fallback&&(!t.exact||t.exact&&n.exact)&&(!t.parent||!t.parent.firstmatch||!t.parent.matched)?(e.onParams(t.params=n.params),t.parent&&(t.parent.matched=!0),t.show()):t.hide(),await tick(),n&&!t.fallback&&(t.childs.size>0&&t.activeChilds.size==0||t.childs.size==0&&t.fallbacks.size>0)){let i=t;for(;i.fallbacks.size==0;)if(i=i.parent,!i)return;i&&i.fallbacks.forEach(s=>s.show());}}};return t.makePattern(e.path),setContext("tinro",t),onMount(()=>t.register()),t.un=o.subscribe(r=>{t.match(r.path);}),t}function f(e,c=!1){return e=e.slice(e.startsWith("/#")?2:0,e.endsWith("/*")?-2:void 0),e.startsWith("/")||(e="/"+e),e==="/"&&(e=""),c&&!e.endsWith("/")&&(e+="/"),e}function h(e,c){e=f(e,!0),c=f(c,!0);let a=[],t={},r=!0,n=e.split("/").map(s=>s.startsWith(":")?(a.push(s.slice(1)),"([^\\/]+)"):s).join("\\/"),i=c.match(new RegExp(`^${n}$`));return i||(r=!1,i=c.match(new RegExp(`^${n}`))),i?(a.forEach((s,w)=>t[s]=i[w+1]),{exact:r,params:t}):null}function v(e,c,a){if(a==="")return e;if(a[0]==="/")return a;let t=i=>i.split("/").filter(s=>s!==""),r=t(e),n=t(c);return "/"+n.map((i,s)=>r[s]).join("/")+"/"+a}function l(e,c,a,t){let r=[c,"data-"+c].reduce((n,i)=>{let s=e.getAttribute(i);return a&&e.removeAttribute(i),s===null?n:s},!1);return !t&&r===""?!0:r||(t||!1)}function d(e){let c=e.split("&").map(a=>a.split("=")).reduce((a,t)=>{let r=t[0];if(!r)return a;let n=t.length>1?t[t.length-1]:!0;return typeof n=="string"&&n.includes(",")&&(n=n.split(",")),a[r]===void 0?a[r]=[n]:a[r].push(n),a},{});return Object.entries(c).reduce((a,t)=>(a[t[0]]=t[1].length>1?t[1]:t[1][0],a),{})}function R(e){throw new Error(e)}var o=j();function j(){let e=window.location.pathname==="srcdoc",c=(r,n)=>{e?window.location.hash=r:history.pushState({},"",r),n(u(e));},{subscribe:a,set:t}=writable(u(e),r=>{window.hashchange=window.onpopstate=i=>t(u(e));let n=C(i=>c(i,t));return i=>{window.hashchange=window.onpopstate=null,n();}});return {subscribe:a,goto:r=>c(r,t),params:L,useHashNavigation:r=>t(u(e=r===void 0?!0:r))}}function u(e){return e?P():{path:window.location.pathname,query:d(window.location.search.slice(1)),hash:window.location.hash.slice(1)}}function P(){let e=String(window.location.hash.slice(1)||"/").match(/^([^?#]+)(?:\?([^#]+))?(?:\#(.+))?$/);return {path:e[1]||"",query:d(e[2]||""),hash:e[3]||""}}function C(e){let c=a=>{let t=a.target.closest("a[href]"),r=t&&l(t,"tinro-ignore");if(!r&&t){let n=t.getAttribute("href").replace(/^\/#/,"");/^\/\/|^[a-zA-Z]+:/.test(n)||(a.preventDefault(),e(n.startsWith("/")?n:t.href.replace(window.location.origin,"")));}};return addEventListener("click",c),()=>removeEventListener("click",c)}function L(){return getContext("tinro").params}

/* node_modules/tinro/cmp/Route.svelte generated by Svelte v3.31.0 */
const get_default_slot_changes = dirty => ({ params: dirty & /*params*/ 2 });
const get_default_slot_context = ctx => ({ params: /*params*/ ctx[1] });

// (23:0) {#if showContent}
function create_if_block(ctx) {
	let current;
	const default_slot_template = /*#slots*/ ctx[7].default;
	const default_slot = create_slot(default_slot_template, ctx, /*$$scope*/ ctx[6], get_default_slot_context);

	return {
		c() {
			if (default_slot) default_slot.c();
		},
		m(target, anchor) {
			if (default_slot) {
				default_slot.m(target, anchor);
			}

			current = true;
		},
		p(ctx, dirty) {
			if (default_slot) {
				if (default_slot.p && dirty & /*$$scope, params*/ 66) {
					update_slot(default_slot, default_slot_template, ctx, /*$$scope*/ ctx[6], dirty, get_default_slot_changes, get_default_slot_context);
				}
			}
		},
		i(local) {
			if (current) return;
			transition_in(default_slot, local);
			current = true;
		},
		o(local) {
			transition_out(default_slot, local);
			current = false;
		},
		d(detaching) {
			if (default_slot) default_slot.d(detaching);
		}
	};
}

function create_fragment(ctx) {
	let if_block_anchor;
	let current;
	let if_block = /*showContent*/ ctx[0] && create_if_block(ctx);

	return {
		c() {
			if (if_block) if_block.c();
			if_block_anchor = empty();
		},
		m(target, anchor) {
			if (if_block) if_block.m(target, anchor);
			insert(target, if_block_anchor, anchor);
			current = true;
		},
		p(ctx, [dirty]) {
			if (/*showContent*/ ctx[0]) {
				if (if_block) {
					if_block.p(ctx, dirty);

					if (dirty & /*showContent*/ 1) {
						transition_in(if_block, 1);
					}
				} else {
					if_block = create_if_block(ctx);
					if_block.c();
					transition_in(if_block, 1);
					if_block.m(if_block_anchor.parentNode, if_block_anchor);
				}
			} else if (if_block) {
				group_outros();

				transition_out(if_block, 1, 1, () => {
					if_block = null;
				});

				check_outros();
			}
		},
		i(local) {
			if (current) return;
			transition_in(if_block);
			current = true;
		},
		o(local) {
			transition_out(if_block);
			current = false;
		},
		d(detaching) {
			if (if_block) if_block.d(detaching);
			if (detaching) detach(if_block_anchor);
		}
	};
}

function instance($$self, $$props, $$invalidate) {
	let { $$slots: slots = {}, $$scope } = $$props;
	let { path = "/*" } = $$props;
	let { fallback = false } = $$props;
	let { redirect = false } = $$props;
	let { firstmatch = false } = $$props;
	let showContent = false;
	let params = {};

	m({
		path,
		fallback,
		redirect,
		firstmatch,
		onShow() {
			$$invalidate(0, showContent = true);
		},
		onHide() {
			$$invalidate(0, showContent = false);
		},
		onParams(newparams) {
			$$invalidate(1, params = newparams);
		}
	});

	$$self.$$set = $$props => {
		if ("path" in $$props) $$invalidate(2, path = $$props.path);
		if ("fallback" in $$props) $$invalidate(3, fallback = $$props.fallback);
		if ("redirect" in $$props) $$invalidate(4, redirect = $$props.redirect);
		if ("firstmatch" in $$props) $$invalidate(5, firstmatch = $$props.firstmatch);
		if ("$$scope" in $$props) $$invalidate(6, $$scope = $$props.$$scope);
	};

	return [showContent, params, path, fallback, redirect, firstmatch, $$scope, slots];
}

class Route extends SvelteComponent {
	constructor(options) {
		super();

		init(this, options, instance, create_fragment, safe_not_equal, {
			path: 2,
			fallback: 3,
			redirect: 4,
			firstmatch: 5
		});
	}
}

export { Route };
