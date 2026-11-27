// Whack-a-snap. Without this script the middle snap is a plain link to
// /random. With it the snap sits up and taunts until the mallet reaches it,
// then ducks and pops up from another hole. It takes three hits: each one
// batters it more, and the third flattens it and opens a random opinion.
// With reduced motion the game still runs, but the CSS drops the sliding
// and the squash, so the snap just appears and disappears.
(() => {
  const snaps = [...document.querySelectorAll(".snap")];
  if (snaps.length === 0) return;
  const still = matchMedia("(prefers-reduced-motion: reduce)");
  const root = document.documentElement;
  const board = document.querySelector(".board");
  let up = snaps.findIndex((s) => !s.classList.contains("down"));
  let timer = 0;
  let paused = false;
  let hit = false; // the snap is dead and the page is on its way to a quote
  const lives = 3;
  let hits = 0;
  const stages = (board.dataset.damage || "").split(" ").filter(Boolean);

  const between = (min, max) => min + Math.random() * (max - min);

  // A snap that is down plays its duck animation once; one that comes up
  // plays the pop. A snap that has never been up starts down without ducking.
  const show = (snap, visible) => {
    snap.classList.remove("pop", "down", "rest", "ouch");
    void snap.offsetWidth; // restart the animation
    snap.classList.add(visible ? "pop" : "down");
  };

  // Exactly one snap is in the tab order: the one that is up, or the one about
  // to come up. Focusing it brings it up at once, so Tab never misses the game.
  const target = (i) => {
    snaps.forEach((snap, j) => {
      if (j === i) {
        snap.removeAttribute("tabindex");
        snap.removeAttribute("aria-hidden");
      } else {
        snap.tabIndex = -1;
        snap.setAttribute("aria-hidden", "true");
      }
    });
  };

  // The snap tires. Every dodge and every miss adds to its fatigue, and so
  // does time spent chasing it; each dodge comes a little later, and once it
  // is worn out it stops running and just sits there taunting. A hit gives it
  // a second wind. So a player who keeps missing still gets through.
  let fatigue = 0;
  let chaseStart = 0;
  const effort = () => fatigue + (chaseStart ? Math.floor((performance.now() - chaseStart) / 1600) : 0);
  const tired = () => effort() >= 6;
  const tire = (by = 1) => {
    fatigue += by;
    chaseStart ||= performance.now();
    board.classList.toggle("tired", tired());
  };

  // With a mouse the snap only moves when the mallet is on it: 0.45 to
  // 0.75 s after the mallet reaches it (later as it tires), it is back in its
  // hole (the 320 ms duck included), then it comes up somewhere else.
  let dodging = false;
  let pointer = null;
  const dodge = () => {
    if (dodging || paused || hit || document.hidden || tired()) return;
    dodging = true;
    const delay = between(450, 750) + effort() * 160 - (still.matches ? 0 : 320);
    tire();
    timer = setTimeout(hop, delay);
  };

  // On a touch screen nothing hovers, so the snap hops on its own instead,
  // every couple of seconds, until it tires.
  const touchOnly = matchMedia("(hover: none)").matches;
  const idle = () => {
    if (!touchOnly || dodging || paused || hit || document.hidden || tired()) return;
    clearTimeout(timer);
    dodging = true;
    timer = setTimeout(() => {
      tire();
      hop();
    }, between(1600, 2600) + effort() * 300);
  };
  // Whether the mallet is on the snap. The browser only reports hover after
  // the mouse moves, so a snap that comes up under a resting mallet checks
  // the last pointer position itself.
  const underMallet = (snap) => {
    if (!pointer) return false;
    const r = snap.getBoundingClientRect();
    return pointer.x >= r.left && pointer.x <= r.right && pointer.y >= r.top && pointer.y <= r.bottom;
  };
  // Stop any dodge in progress. If it was caught between holes, bring it up.
  const cancel = () => {
    clearTimeout(timer);
    dodging = false;
    if (!hit && snaps[up].classList.contains("down")) show(snaps[up], true);
  };

  // Drop the snap, wait a moment, then pop it up from another hole.
  const hop = () => {
    let next = up;
    while (next === up) next = Math.floor(Math.random() * snaps.length);
    dodging = false;
    show(snaps[up], false);
    up = next;
    target(up);
    timer = setTimeout(() => {
      if (hit || paused) return;
      show(snaps[up], true);
      look();
      // It may come up right under the mallet: check once it is out of the hole.
      const snap = snaps[up];
      setTimeout(() => {
        if (snap === snaps[up] && underMallet(snap)) dodge();
      }, still.matches ? 0 : 220);
      idle();
    }, between(450, 650));
  };

  // Sound, made on the spot with Web Audio: a hollow bonk and a cardboard
  // crunch for a hit, a dull thud for a miss. It only ever plays in answer
  // to a click, tap or key press.
  let audio = null;
  const sound = (kind) => {
    const isHit = kind !== "miss";
    const heavy = kind === "kill";
    try {
      audio ??= new (window.AudioContext || window.webkitAudioContext)();
      if (audio.state === "suspended") audio.resume();
      const t = audio.currentTime;
      const out = audio.createGain();
      out.gain.value = heavy ? 0.5 : isHit ? 0.4 : 0.25;
      out.connect(audio.destination);

      const body = audio.createOscillator();
      const bodyGain = audio.createGain();
      body.type = "triangle";
      // Each hit lands a little lower, so the last one sounds final.
      const pitch = heavy ? 170 : isHit ? 260 - hits * 30 : 130;
      body.frequency.setValueAtTime(pitch, t);
      body.frequency.exponentialRampToValueAtTime(heavy ? 42 : isHit ? 60 : 45, t + (heavy ? 0.24 : 0.16));
      bodyGain.gain.setValueAtTime(1, t);
      bodyGain.gain.exponentialRampToValueAtTime(0.001, t + (heavy ? 0.3 : isHit ? 0.2 : 0.14));
      body.connect(bodyGain).connect(out);
      body.start(t);
      body.stop(t + 0.26);

      const length = Math.floor(audio.sampleRate * (heavy ? 0.22 : isHit ? 0.09 : 0.06));
      const noise = audio.createBuffer(1, length, audio.sampleRate);
      const data = noise.getChannelData(0);
      for (let i = 0; i < length; i++) data[i] = (Math.random() * 2 - 1) * (1 - i / length) ** 2;
      const crunch = audio.createBufferSource();
      const filter = audio.createBiquadFilter();
      const crunchGain = audio.createGain();
      crunch.buffer = noise;
      filter.type = isHit ? "bandpass" : "lowpass";
      filter.frequency.value = isHit ? 1100 : 400;
      crunchGain.gain.value = heavy ? 1.1 : isHit ? 0.7 : 0.5;
      crunch.connect(filter).connect(crunchGain).connect(out);
      crunch.start(t);
    } catch {
      // No audio is fine.
    }
  };

  const label = () => {
    const left = lives - hits;
    for (const snap of snaps) {
      snap.setAttribute("aria-label", `Whack a snap: ${left} ${left === 1 ? "hit" : "hits"} left to open a random opinion`);
    }
  };
  label();

  // Every snap shows the same battered box, wherever it comes up.
  const batter = () => {
    const src = stages[hits - 1];
    if (!src) return;
    for (const snap of snaps) {
      snap.querySelector("source")?.setAttribute("srcset", src);
      snap.querySelector(".box").src = src;
    }
  };
  // Fetch the battered boxes ahead of the first hit.
  const preload = () => {
    for (const src of stages) new Image().src = src;
  };
  addEventListener("pointermove", preload, { once: true });
  addEventListener("keydown", preload, { once: true });

  const thud = () => {
    board.classList.remove("thud");
    void board.offsetWidth;
    board.classList.add("thud");
  };

  // A hit. The first two batter it; the third flattens it and opens a random
  // opinion. A snap that is up flinches and, under the mallet, runs off at
  // once; one that is already sliding down takes the damage and keeps going.
  const status = document.getElementById("whack-status");
  const say = (text) => {
    if (status) status.textContent = text;
  };

  const whack = (snap, byMouse) => {
    if (hit) return;
    hits++;
    thud();
    // A hit gives it a second wind.
    fatigue = Math.max(0, fatigue - 2);
    chaseStart = performance.now();
    board.classList.remove("tired");
    if (hits >= lives) {
      hit = true;
      clearTimeout(timer);
      dodging = false;
      snap.classList.remove("pop", "down", "rest", "ouch");
      snap.classList.add("whacked");
      for (const other of snaps) if (other !== snap && !other.classList.contains("down")) show(other, false);
      sound("kill");
      say("Got it. Opening a random opinion.");
      setTimeout(() => location.assign(snap.href), still.matches ? 300 : 650);
      return;
    }
    sound("hit");
    batter();
    label();
    const left = lives - hits;
    say(`Ouch. ${left} ${left === 1 ? "hit" : "hits"} left.`);
    if (snap.classList.contains("down")) return;
    snap.classList.remove("pop", "ouch");
    void snap.offsetWidth;
    snap.classList.add("ouch");
    if (byMouse && !paused && !tired()) {
      clearTimeout(timer);
      dodging = true;
      timer = setTimeout(hop, still.matches ? 150 : 260);
    } else {
      dodging = false;
      idle();
    }
  };

  // Where the head lands decides the hit: any snap whose drawn box shows
  // above its hole's rim at that point is hit, whether it is popping up,
  // flinching or sliding down. The plate's transparent margin does not count.
  const struck = (x, y) => {
    for (const snap of snaps) {
      if (snap.classList.contains("whacked")) continue;
      const box = snap.querySelector(".box").getBoundingClientRect();
      const rim = snap.closest(".window").getBoundingClientRect().bottom;
      const top = box.top + box.height * 0.06;
      const bottom = Math.min(box.top + box.height * 0.92, rim);
      const left = box.left + box.width * 0.075;
      const right = box.right - box.width * 0.075;
      if (top < bottom && x >= left && x <= right && y >= top && y <= bottom) return snap;
    }
    return null;
  };

  for (const snap of snaps) {
    snap.addEventListener("pointerenter", (e) => {
      if (e.pointerType === "mouse" && snap === snaps[up]) dodge();
    });
    // Keyboard and touch whack on click; the mouse whacks on press, below,
    // so the click that follows a mouse press is not a second hit.
    snap.addEventListener("click", (e) => {
      if (snap.classList.contains("down") || hit || performance.now() - swungAt < 1000) {
        e.preventDefault();
        return;
      }
      if (e.metaKey || e.ctrlKey || e.shiftKey || e.button > 0) return;
      e.preventDefault();
      whack(snap, false);
    });
    // A keyboard user who is about to whack it should not have it run away.
    snap.addEventListener("focus", () => {
      paused = true;
      cancel();
    });
    snap.addEventListener("blur", () => {
      paused = false;
      if (underMallet(snap)) dodge();
      else idle();
    });
  }

  // The snap that is up keeps its eyes on the mallet.
  let frame = 0;
  const look = () => {
    frame = 0;
    const snap = snaps[up];
    if (!pointer || !snap) return;
    const eyes = snap.querySelector(".eyes").getBoundingClientRect();
    const dx = pointer.x - (eyes.left + eyes.width * 0.35);
    const dy = pointer.y - (eyes.top + eyes.height * 0.53);
    const clamp = (v) => Math.max(-1, Math.min(1, v));
    snap.style.setProperty("--ex", clamp(dx / (eyes.width * 0.8)).toFixed(2));
    snap.style.setProperty("--ey", clamp(dy / (eyes.height * 0.8)).toFixed(2));
  };

  // With a mouse the pointer is a drawn mallet, held at the end of its handle,
  // so a swing brings the head down onto the pointer. Without this script the
  // CSS falls back to a plain mallet cursor.
  const mallet = document.querySelector(".mallet");
  const fine = matchMedia("(hover: hover) and (pointer: fine)");
  if (mallet && fine.matches) root.classList.add("js-mallet");
  let moved = 0;
  const follow = () => {
    moved = 0;
    if (!pointer) return;
    mallet.style.setProperty("--mx", `${pointer.x}px`);
    mallet.style.setProperty("--my", `${pointer.y}px`);
  };
  // Within reach of the theme picker the mallet is put away, so its head
  // (drawn above the pointer) never covers the button.
  const picker = document.querySelector(".theme");
  const nearPicker = (x, y) => {
    if (!picker || picker.hidden) return false;
    const r = picker.getBoundingClientRect();
    return Math.hypot(x - (r.left + r.width / 2), y - (r.top + r.height / 2)) < 120;
  };
  addEventListener("pointermove", (e) => {
    if (e.pointerType !== "mouse") return;
    pointer = { x: e.clientX, y: e.clientY };
    root.classList.remove("away");
    root.classList.toggle("near-ui", nearPicker(e.clientX, e.clientY));
    if (!frame) frame = requestAnimationFrame(look);
    if (mallet && !moved) moved = requestAnimationFrame(follow);
  });
  document.addEventListener("mouseout", (e) => {
    if (!e.relatedTarget) root.classList.add("away");
  });

  // Press: the head comes down, and lands 70 ms later. That is the moment
  // of the hit, the squash and the sound.
  let swungAt = -Infinity;
  addEventListener("pointerdown", (e) => {
    if (e.pointerType !== "mouse" || e.button !== 0) return;
    // Over the footer band the mallet is put away: no swing, no sound.
    if (e.target.closest?.(".report, .theme") || root.classList.contains("near-ui")) return;
    root.classList.remove("swing");
    void root.offsetWidth;
    root.classList.add("swing");
    if (e.metaKey || e.ctrlKey || e.shiftKey) return; // open the link elsewhere
    swungAt = performance.now();
    if (e.target.closest?.(".snap")) e.preventDefault();
    setTimeout(() => {
      if (hit) return;
      const at = pointer ?? { x: e.clientX, y: e.clientY };
      const snap = struck(at.x, at.y);
      if (snap) whack(snap, true);
      else if (e.target.closest?.(".board")) {
        // A miss near the holes: a dull thud, and the snap tires a little.
        sound("miss");
        tire();
      }
    }, still.matches ? 0 : 70);
  });
  const lift = () => {
    const left = Math.max(0, 180 - (performance.now() - swungAt));
    setTimeout(() => root.classList.remove("swing"), left);
  };
  addEventListener("pointerup", lift);
  addEventListener("pointercancel", lift);

  document.addEventListener("visibilitychange", () => {
    if (document.hidden) cancel();
  });
  // Coming back from the quote with the Back button would restore the
  // flattened box mid-game; start a fresh game instead.
  addEventListener("pageshow", (e) => {
    if (e.persisted) location.reload();
  });
  idle();
})();
