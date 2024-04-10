package icu.aimerny.elix;

import icu.aimerny.elix.registry.ModItem;
import net.fabricmc.api.ModInitializer;

public class Elix implements ModInitializer {

    @Override
    public void onInitialize() {
        ModItem.init();
    }
}
